#!/usr/bin/python
#-*- encoding:utf-8 -*-
#建立一个窗口

import gtk
class PyApp(gtk.Window):
    def __init__(self):
        super(PyApp, self).__init__()
        self.set_title("PyGtk")
        self.set_size_request(250, 150)
        self.set_position(gtk.WIN_POS_CENTER)

        self.connect("destroy", gtk.main_quit)

        self.show()

    def main(self):
        gtk.main()

print __name__
if __name__ == "__main__":
    pyapp = PyApp()
    pyapp.main()
