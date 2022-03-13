#!/usr/bin/env zx

import 'fs';
import 'path';

// configurations
$.verbose = false

// utility functions
const groupBy = (array, toKey) => array.reduce((obj, current) => {
  const key = toKey(current);
  (obj[key] || (obj[key] = [])).push(current);
  return obj;
}, {});

const first = (array) => array[0]
const last = (array) => array[array.length - 1]
const isNotEmpty = (text) => text

const difference = (l, r) => l.filter((e) => !r.includes(e))

// domain specific functions
const toLinked = async (p) => last((await $`ls -l ${p}`).stdout.split(' ')).trim()
const toLinkedDir = async (p) => path.dirname(await toLinked(p))

const toObject = (dir, paths) => {
  const obj = {}
  obj['dir'] = dir
  obj['files'] = paths.map((p) => path.basename(p))
  return obj
}

const appendSource = async (d) => {
  const sourceDir = await toLinkedDir(path.resolve(d['dir'], first(d['files'])));
  const sourceFiles = fs.readdirSync(path.resolve(d['dir'], sourceDir));

  d['sourceDir'] = sourceDir
  d['sourceFiles'] = sourceFiles
  d['nonLinkedSourceFiles'] = difference(sourceFiles, d['files'])

  return d
}

// process input
const args = argv['_'].slice(1)

if (args.length != 1) {
  console.log('usage: ./ccc.mjs [target path]')
  process.exit()
}

const target = args[0]

// execute
const found = await $`find ${target} -type l -name '*.json'`

const symlinkPaths = found.stdout.split(/\n/).filter(isNotEmpty)
const symlinkDirs = groupBy(symlinkPaths, (p) => path.dirname(p))

const objects = Object.entries(symlinkDirs)
  .map(([dir, paths]) => toObject(dir, paths))
  .map(appendSource)

console.log(JSON.stringify(await Promise.all(objects), null, 2))
