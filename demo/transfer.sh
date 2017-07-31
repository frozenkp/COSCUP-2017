#!/bin/sh

path=$1
file=$2

cp title $file

ls -al $path | awk '
BEGIN{
  a=0;
}
{
  if(a > 2){
    system("go run imgToData.go " path "/" $9 " >> " file);
  }
  a++;
}
' file="$file" path="$path"
