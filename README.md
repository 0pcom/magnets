```
cd voice_messages ; echo '<!doctype html><html lang=en><head></head><body style="background-color:black;color:white;"><br>' | tee index.html ; for i in $(ls --ignore index.html) ; do  echo "<figure><figcaption>${i}</figcaption><audio controls src=\"${i}\"><a href=\"${i}\">Download audio</a></audio></figure>" | tee -a index.html ; done ; echo "</body></html>" | tee -a index.html ; cd ..
```
