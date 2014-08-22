#!/usr/bin/env python3

import sys


DICT_DIR = '/home/dave/dict/dict'
WORD_LIST_DIR = '/usr/share/dict/american-english'


def checkWordList(word):
    """
    check wordlist if word in wordlist return Ture
    """
    flag = None
    with open(WORD_LIST_DIR) as f:
        for line in f:
            if line.strip().lower() == word:
                # convert the word in list to lowercase
                flag = True
    return flag


def searchDictList(word):
    """
    search dictList.
    if not found return None
    else return the found line
    """
    text = None
    with open(DICT_DIR) as f:
        for line in f:
            if line.split(':')[0] == word:
                text = line
    return text


def sortList(dir=DICT_DIR):
    """
    sort the dictList default value is DICT_DIR
    and remove the space
    """
    with open(DICT_DIR) as f:
        lines = [line for line in f if line.strip()]  # remove space
        lines.sort()
    with open(DICT_DIR, 'w') as f:
        f.writelines(lines)


def getWord(word):
    """
    add the user input
    """
    list = ['英音', '美音', 'vt.', 'vi.', 'adj.', 'adv.', 'n.', 'other. ']
    print(''+word)
    content = str()
    for _ in list:
        text = input(_).strip()
        if text:
            content += ':' + _ + text
    line = word + content
    return line


def addWord(line):
    with open(DICT_DIR, 'a', encoding='utf8') as f:
        f.write(line+'\n')
    sortList()  # call sortList function
    print('Add successfully!')


def countWords():
    """
    traversal
    """
    nu = 0
    for line in open(DICT_DIR):
        nu += 1
    print()
    print('(^_^): already have {} words!'.format(nu))
    print()


def modifyWord(word):
    if checkWordList(word) and searchDictList(word):
        pass
    else:
        print("\n(~_~): Are you sure {0:-^8} is a word ".format(word))
        print("or {0:-^8} already in dict list".format(word))
        sys.exit()
    if input('\nAre you sure to modify {}?(yes/no):\n'.format(word)) == 'yes':
        with open(DICT_DIR) as f:
            old = f.readlines()
        newLine = getWord(word)
        for no in range(len(old)):
            if old[no].split(':')[0] == word:
                old[no] = newLine + '\n'
        with open(DICT_DIR, 'w') as f:
            f.writelines(old)
        print('Add successfully!')


def deleteWord(word):
    if checkWordList(word) and searchDictList(word):
        pass
    else:
        print("\n(~_~): Are you sure {0:-^8} is a word ".format(word))
        print("or {0:-^8} already in dict list".format(word))
        sys.exit()
    if input('\nAre you sure to delete {}?(yes/no):\n'.format(word)) == 'yes':
        with open(DICT_DIR) as f:
            old = f.readlines()
        for no in range(len(old)):
            if old[no].split(':')[0] == word:
                old[no] = '\n'
        with open(DICT_DIR, 'w') as f:
            f.writelines(old)
        print('Delete successfully!')


def printWord(line):
    for _ in line.split(':'):
        print(_)


def main():
    if not words:  # if wordlist is None exit
        sys.exit('(~_~):no word to dict...!')
    for word in words:
        word = word.lower()  # all input convert to lowercase
        """
        if not in wordlist sys.exit()
        """
        flag = checkWordList(word)
        if not flag:
            print()
            print("(~_~): Are you sure {:-^8} is a word?".format(word))
            print()
            sys.exit()

        """
        search wordlist
        """
        line = searchDictList(word)
        if line:  # if find the word record print it.
            printWord(line)
        else:     # else interactive ask user wheather add the word
            print('(#_#): {:-^8} not in my dictionary!'.format(word))
            print(' ')
            yes = {'y', 'yes'}
            if input('add {}. y/n: '.format(word)) in yes:
                line = getWord(word)
                addWord(line)


if __name__ == "__main__":
    words = None
    options = {'--count', '-m', '-r'}
    if len(sys.argv) == 1:
        print('(*_*)Usage:{} [--options] [word1] [word2] ...'
              .format(sys.argv[0]))
        sys.exit()
    else:
        cmd = sys.argv[1:]
    if cmd[0] == '--count':
        countWords()
    elif cmd[0] == '-m':
        words = cmd[1:]
        for word in words:
            modifyWord(word)
    elif cmd[0] == '-d':
        words = sys.argv[2:]
        for word in words:
            deleteWord(word)
    else:
        words = cmd
        main()
