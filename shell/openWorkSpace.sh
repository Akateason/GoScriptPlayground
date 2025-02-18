#!/bin/bash

find . -name "*.xcworkspace" -print -quit | while read workspace; do  
  open "$workspace"  
  break
done
