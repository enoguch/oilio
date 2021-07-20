---
title: "_Index"
date: 2021-07-12
tags: []
draft: false
---


[![CI](https://github.com/enoguch/oilio/actions/workflows/main.yml/badge.svg)](https://github.com/enoguch/oilio/actions/workflows/main.yml)
[![Coverage Status](https://coveralls.io/repos/github/enoguch/oilio/badge.svg?branch=main)](https://coveralls.io/github/enoguch/oilio?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/enoguch/oilio)](https://goreportcard.com/report/github.com/enoguch/oilio)
[![license](http://img.shields.io/badge/license-CC0-green.svg)](https://github.com/enoguch/oilio/blob/main/LICENSE)

[![Docker](https://img.shields.io/badge/Docker-ghcr.io%2Fenoguch%2Foilio%3A1.0.0-green?logo=docker)](https://hub.docker.com/repository/docker/enoguch/oilio)
[![DOI](https://zenodo.org/badge/370310415.svg)](https://zenodo.org/badge/latestdoi/370310415)


# oilio
wcコマンドを拡張したソフトウェア

## 概要
与えられたファイルもしくは与えられたディレクトリの直下のファイル群の
行数、文字数、バイト数をカウントする．
* .gitignoreファイルを参照できる
* アーカイブファイルも扱える(jar, tar, tazなど)

# 使用方法
oilio [CLI_MODE_OPTION] [FILEs...|DIRs...]
* コマンドオプション
  * -b, --byte         バイト数表示
  * -c, --character    文字数表示
  * -l, --line         行数表示
  * -n, --no-ignore    .gitignore内のファイルを見ない
  * -h, --help         ヘルプ表示

# アイコン
![whale-animal-publicdomainvectors](https://user-images.githubusercontent.com/84704334/119367448-ed089b00-bcec-11eb-850e-21628810b8d8.png)

アイコンは[freesvg.org](https://freesvg.org/whale)より取得

