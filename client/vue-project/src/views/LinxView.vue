<template>
    <div id="linx">
        
        <br>
      <h1>This is the Linx Page</h1>
      <linximporter @load="LoadHandler"></linximporter>
        <div class="grid-container"  > <!-- v-for state === unknown word, class = unknownword  -->
          <!-- <div :style="{color : activeColor}" :id="`container-${index}`">{{ word }}   XXX possibly add a check in here to change the colour accordingly-->
          <linxreader v-for="[word, state], index in wordMap" :id="`wordthing${index}`" :renderWord="word" :learned="state" :class="linxreadermethod(renderWord)" @click="linxreadermethod(this.renderWord)" > </linxreader>
            <!--TODO update wordthing dynamically -->
        </div> <!-- TODO TODO replace wordMap with simply a list of words in the text file, then iterate through each word, make a call back to the server, have the server return a JSON object for each word. Before all of this create a default JSON object to initalize each word, each local JSON object should then be passed back to server for storage.-->
    
      </div>
<!-- TODO convert from using struct from Linx component to using props in LinxReader component-->


  </template>
 //XXX Might be better to iterate over the map in the method beforehand
<script>
    import { stringifyExpression } from "@vue/compiler-core";
import linximporter from "../components/Linx.vue"
    import linxreader from "../components/LinxReader.vue"
    import axios from 'axios';
import { render } from "vue";
    export default{
      mounted : function() {
          this.HandleTextRender();
        },
        data () {
          
          //XXX they've probably done it by simply making different classes for every kind of interaction and then simply switching the interaction on event

          return {
          activeColor: 'red',
          mySentence : {},
          rawhtml : '<span style="color: red">This should be red.</span>',
          wordMap : Map,
          state : "",
          word : "",

          renderState: null,
          learned: null
          }//one sentence class holds multiple sentenceitems
          //TODO keeping wordState in data return might not be necessary. Keep an array of both the word and the wordState in data return and simply modify the state within a method to allow for freezing. It should still update the rendering when one of the elements of the array changes.
          
        },
        components: {
        linximporter,
        linxreader
        }
        ,
        methods: {
          LoadHandler(textMap, sentence) {
            this.wordMap = textMap;
            this.mySentence = sentence;
          },
          HandleTextRender(){

          axios({ method: "GET", params: {QueryWord: "prima"}, responseType: 'json', url: " http://localhost:8030/fetchWord"}).then(function (response) {
            console.log(response.data);
            //console.log(JSON.stringify(JSON.parse(response), null, 2));
          });
      


        },
        LinxReaderLoaded : function(event) {
          console.log(event.target.id);
          //Make call when this is called to call back to server to check for user data and dynamically update the CSS styling from there

        },
        linxreadermethod: function(renderWord) {
            var classA = '';
            console.log(renderWord);
            if (this.learned == this.mySentence.sentenceitem.unknownword) {
              classA = 'wordthingKnown';
            }
            return classA;
          }
      },
        computed: {
          linxreaderclass: function(word) {
            var classA = '';
            console.log(this.renderWord);
            if (this.renderWord == "CREATE") {
              classA = 'wordthingKnown';
            }
            return classA;
          }
        }

        
    }
  
  

</script>

<style>
.grid-container {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  background-color: #10436d;
  padding: 10px;
  grid-auto-rows: minmax(100px, auto);
}
.wordthingKnown {
  display:inline;
  background-color: rgba(221, 10, 10, 0.8);
  border: 1px solid rgba(25, 0, 255, 0.8);
  padding: 20px;
  font-size: 30px;
  text-align: center;
  font: bold;
}
.wordthingUnknown {
  display:inline;
  background-color: rgba(0, 255, 179, 0.8);
  border: 1px solid rgba(0, 0, 0, 0.8);
  padding: 20px;
  font-size: 30px;
  text-align: center;
}

</style>