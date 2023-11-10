# PersonalWebsiteJS/Sharp-ReaderJS

This is a reimplementation of my Sharp-Reader project using Vue.JS, PostgreSQL, and Go.
Partially chosen as a learning experience but also to make the application more portable by moving away from .Net.

Currently the project is written as a SPA application which allows a user to read texts (currently Italian ones), click on words to get their definitions, and add them to a spaced repetition learning system similar to that of Anki.

Development will be resumed after the stemmer is finished.
Currently dictionaries aren't made available here so as to not step on anyone's toes, scripts for getting dictionaries and setting them up in a local database will be made available in the future.
## Goals

- Write a clean re-implementation of my other .Net-based project using Vue.JS, PostgreSQL, and Go.
- Implement an independently-developed version of the Snowball stemming algorithm and make it easy to add support for other languages.
- Provide a way for users to have an easy-to-set-up offline version of a website like Lingq.
- Provide options for users to use either GUI- or TUI based systems.
