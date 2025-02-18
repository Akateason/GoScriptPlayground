#!/bin/bash

find . -name "*.xcworkspace" ! -path "*/project.xcworkspace" -print -quit | while read workspace; do  
  echo "即将打开😁 $workspace"
  open "$workspace"
  break
done
