function ghostMonitor() {
  if [[ $(pgrep GoHostMonitor) != null ]]; then
      sudo GoHostMonitor
  fi
  ghostMonitor
}
ghostMonitor