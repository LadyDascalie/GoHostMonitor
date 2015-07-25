#!/bin/bash
killall mDNSResponder \
&& osascript -e  'display  notification  "Hosts  changed:  mDNSResponder  restarted"  with  title  "GoHostMonitor"' \
sh ./ghostMonitor.sh