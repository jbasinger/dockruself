# Dockurself runner

## Intro

Testing actions locally is really difficult. Being able to run locally would be ultra powerful.
Lets make one

## Components

* YAML reader
  * Need to be able to read actions into data structure
* Docker Runner
  * go docker lib


## Dependencies
  * YAML Reader
  * Docker lib
  * Git lib

## Things to think about

* Would be nice to convert yaml specs to our struct, and be able to do so with other 
  runner things as well, like GitLab or whatever

* Start simple. Take a single job, a runs on, only shell step to start.

* Next, work on being able to pull actions and run them
  * Will likely require a git lib

* Levels of things like workflows -> jobs -> steps
