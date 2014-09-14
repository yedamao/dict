#!/usr/bin/env python
# -*- coding: utf-8 -*-

import MySQLdb

DICT_DIR = '/usr/share/dict/american-english'

db = MySQLdb.connect("localhost", user="root", passwd="telecom",
                     db="dictionary", charset="utf8")
cursor = db.cursor()
cmd = "insert into word_list (word, length) values ('"

with open(DICT_DIR) as f:
    # count = 0
    for line in f:
        # count += 1
        word = line.strip()
        sql = cmd + word + "',length('" + word + "')" + ");"
        print sql
        # print count
        print ("\n")

        try:
            cursor.execute(sql)
            db.commit()
            print "succssful"
        except MySQLdb.Error, e:
            db.rollback()
            print "Mysql Error %d: %s" % (e.args[0], e.args[1])
            print "error"

db.close()
