#!/bin/bash
godepgraph -nostdlib -novendor $1 | dot -Tpng -o $2