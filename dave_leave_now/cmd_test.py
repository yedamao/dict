#!/usr/bin/env python

import cmd

DICT_DIR = '/home/dave/dict/dictionary'


def searchDictList(word):
    """
    search dictList.
    if not found return None
    else return the found line
    """
    text = None
    with open(DICT_DIR) as f:
        for line in f:
            if line.split(':')[0].lower() == word:
                text = line
    return text


def printWord(line):
    for _ in line.split(':'):
        print(_)


class HelloWorld(cmd.Cmd):
    def __init__(self):
        cmd.Cmd.__init__(self)
        self.intro = """Open Source Dictionary 1.0"""
        self.prompt = "(OSD1.0:>)"

    def do_dict(self, word):
        line = searchDictList(word)
        if line:
            printWord(line)

    def do_greet(self, person):
        """
        greet [person]
        Greet the named person
        """
        if person:
            print ("hi ,", person)
        else:
            print ('hello')

    def help_greet(self):
        print('\n'.join(['greet [person]',
                         'Greet the named person',
                         ]))

    def do_foo(self, line):
        """
        foolish
        """
        print ("bar")

    def do_EOF(self, line):
        return True

if __name__ == '__main__':
    HelloWorld().cmdloop()
