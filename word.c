/*
 * =====================================================================================
 *
 *       Filename:  word.c
 *
 *    Description:  word
 *
 *        Version:  1.0
 *        Created:  2014年10月15日 21时59分11秒
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Dave (), logindaveye@gmail.com
 *   Organization:  
 *
 * =====================================================================================
 */
#include <stdlib.h>
#include <stdio.h>

char word[] = "aaa";
char *p = word;

void print(void)
{
    printf("%s\n", word);
    for(int i = 0; i < 3; i++)
    {
        if (word[i] >= 122)
        {
            word[i] = 'a';
        }
    }
}

void loop(int n)
{
    for(int i = 0; i < 26; i++)
    {
        print();
        (*(p+n))++;
    }
}

int main(void)
{
    /* for(int i = 0; i < 26; i++) */
    /* { */
    /*     print(); */
    /*     (*p)++; */
    /* } */
    loop(0);
    printf("%s\n", word);
}
