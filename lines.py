#!/usr/bin/env python
"""Test script to produce some lines for dots to consume"""

import sys
import time

while True:
    print('line')
    sys.stdout.flush()
    time.sleep(1)
