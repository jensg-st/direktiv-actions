#!/bin/bash

git tag -d test
git push --delete origin test

git tag test
git push origin test
