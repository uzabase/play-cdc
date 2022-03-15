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

const appendSource = async (d, target) => {
  const sourceDir = await toLinkedDir(path.resolve(d['dir'], first(d['files'])));
  const sourceFiles = fs.readdirSync(path.resolve(d['dir'], sourceDir));

  d['relativeDir'] = d['dir'].replace(target, '')
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
const found = (await $`find ${target} -type l -name '*.json'`).stdout

const symlinkPaths = found.split(/\n/).filter(isNotEmpty)
const symlinkDirs = groupBy(symlinkPaths, (p) => path.dirname(p))

const objects = await Promise.all(
  Object.entries(symlinkDirs)
    .map(([dir, paths]) => toObject(dir, paths))
    .map((d) => appendSource(d, target))
  )

const relativeDirs = objects.filter((o) => o.nonLinkedSourceFiles.length > 0).map((o) => o.relativeDir)
console.log(relativeDirs)
const relativeDir = await question('Choose dir for consumer from the above list (use tab): ', {
  choices: relativeDirs
})
const dir = target + relativeDir

cd(dir)

const targetObject = objects.find((o) => o.dir === dir)

console.log(targetObject.nonLinkedSourceFiles)
const file = await question('Choose file to simlink from the above list (use tab): ', {
  choices: targetObject.nonLinkedSourceFiles
})

const fileRelativePath = `${targetObject.sourceDir}/${file}`

const {exitCode} = await $`ln -s ${fileRelativePath}`
console.log('')

if (exitCode === 0) {
  console.log('Linked:')
  console.log((await $`ls -l ${dir}/${file}`).stdout)
} else {
  console.log('Failed.')
}
