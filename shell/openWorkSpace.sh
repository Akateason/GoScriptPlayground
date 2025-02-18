#!/bin/bash

find . -name "*.xcworkspace" -print -quit | while read workspace; do  
  echo "$workspace"  
  break
done
