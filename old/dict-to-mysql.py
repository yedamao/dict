#!/usr/bin/env python
# -*- coding: utf-8 -*-

import MySQLdb

DICT_DIR = '/home/dave/dict/dict'


def insertLine(line):
    word = line.split(':')[0]
    print word
    text = line.split(':')[1:]
    eng = us = vt = vi = adj = adv = n = other = ''
    for content in text:
        if content.find('vt') != -1:
            vt = content
            print vt
        elif content.find('vi') != -1:
            vi = content
            print vi
        elif content.find('adj') != -1:
            adj = content
            print adj
        elif content.find('adv') != -1:
            adv = content
            print adv
        elif content.find('n') != -1:
            n = content
            print n
        elif content.find('other') != -1:
            other = content
            print other
        elif content.find("英") != -1:
            eng = content
            print eng
        elif content.find('美') != -1:
            us = content
            print us

    sql = """INSERT INTO vocabulary ( word,
             `英`, `美`, vt, vi, adj, adv,
             n, other)
             values ("""
    query = '"' + word + '","' + eng + '","' + us + '","' + vt + '","' + vi + '","' + adj + '","' + adv + '","' + n + '","' + other + '")'

    sql = sql + query
    print sql

    try:
        cursor.execute(sql)
        db.commit()
    except MySQLdb.Error, e:
        db.rollback()
        print "Mysql Error %d: %s" % (e.args[0], e.args[1])
        print "error"
    except UnicodeError:
        pass

with open(DICT_DIR) as f:
    db = MySQLdb.connect(host="localhost", user="root", passwd="telecom", db="dict", charset='utf8')
    cursor = db.cursor()

    for line in f:
        insertLine(line)

    db.close()
