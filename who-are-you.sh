#!/bin/bash
curl https://01.gritlab.ax/assets/superhero/all.json | jq '[map(select(.id == 70))] | .[] | .[].name'