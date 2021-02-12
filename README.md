![ACM-HEADER](https://user-images.githubusercontent.com/14032427/92643737-e6252e00-f2ff-11ea-8a51-1f1b69caba9f.png)

<h1 align="center"> ACM-VIT cli </h1>

<p align="center"> 
The official terminal application of acmvit
</p>

<p>
  <a href="https://acmvit.in/" target="_blank">
    <img alt="made-by-acm" src="https://img.shields.io/badge/MADE%20BY-ACM%20VIT-blue?style=for-the-badge" />
  </a>
    <!-- Uncomment the below line to add the license badge. Make sure the right license badge is reflected. -->
    <!-- <img alt="license" src="https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge" /> -->
    <!-- forks/stars/tech stack in the form of badges from https://shields.io/ -->
</p>

---

This project currently exists to fullfil these needs

1) create a project such that :
    - all resource link would no longer be on whatsapp descriptions and lost in discord pinned but organized in one place easily accessible from the cli 
2) create meetings which would be stored in the db + the acmvit google calender :
    - Why ? so that it makes it easier to know what other of different departments are working on and what active projects are working towards so that a member can contribute
    

---

## Installation

1) From npm 
    *this method has know issues i.e permission error by npm its inertied from the go-npm package we use, tho this method is the best if it works cause you will recieve all updates automatically and its installed in your bin by default allowing you to access the cli from anywhere*

   <a href="https://www.npmjs.com/package/acmvit?activeTab=readme" target="_blank">
        The package on npm 
  </a>

  2) From binaries 
    *foolproof fully tested no errors*
     <a href="https://github.com/ACM-VIT/acmInternal-cli/releases">
        All published releases's page on github
     </a>

     choose amd64 for linux
            .exe for windows
            and generally 386 works for debian rpm based distributions
<!-- Add one screenshot of your project (max height: 1000px, max size: 1mb) 'if applicable' under assets folder in root of your project ![sceenshot](assets/<name of image>) -->
<!-- if your project has multiple pictures , merge them into one image using a tool similar to figma -->

---

## Usage
<!-- How To, Features, Installation etc. as subheadings in this section. example-->

### gotachas
1) set cli password from the our official android or react-native app
2) if you have downloaded the binaries manually its advisable to add it to a folder in your path so that you can access the cli from anywhere on your computer
3) its also advisable if you want to put a command alias to change acmvit to another command like acm though this is hard to do in windows

### Commands usage
acmvit help
    - to get all commands
acmvit help (commandName)
    - to get longer usage info on a specific command


## Contributing 

## code
 this whole project is built using golang;

 the main package to get the cli stuff working is known as cobra.

 this project extensively makes calls to the 
 <a href="https://github.com/ACM-VIT/acmInternalBackend">acm internal backend</a>

 the main supporiting internal package is cli-core

 ## building and deploying
   the deploy and build is done using goreleaser + go-npm 

   goreleaser has been made into a github action 

   therfore once you create a new github tag:
   git tag -a v1.1.0 -m "feat: new feature"
   the github action automatically runs

   *important*
   the version in the package.json must be the same as the new tag you plan to create for goreleaser to work







## Authors

**Authors:** <a href="https://github.com/madrigal1">Madrigal1</a>
**Contributors:** <!-- Generate contributors list using this link - https://contributors-img.web.app/preview -->
