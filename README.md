# README

I needed a tool that would watch my `/etc/` folder so as to reset `mDNSResponder` whenever I made change in my `hosts` file. I thought it would make for a good pet project to start learning Go !

Feel free to do whatever you want with the code !

## USAGE

***This "tool" will only work on OSX. It SHOULD also work on linux, but don't take my word for it!***

The best way to use this (admitedly hacky) tool, is as follows:

`cd` into the GoHostMonitor directory, and from there:

`sudo sh ghostMonitor.sh`

This will start the program, and make sure it relaunches after it is done.
Of course, since this works on the principle of an infinite loop, `Ctrl+C` won't exit the script, as it will immediately relaunch. so use : `Ctrl+Z` to suspend it instead !