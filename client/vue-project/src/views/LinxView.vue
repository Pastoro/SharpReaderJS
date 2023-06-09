<template>
    <div id="linx">
        
        <br>
      <h1>This is the Linx Page</h1>
      <linximporter @load="LoadHandler"></linximporter>
        <div class="content" v-for="[word, state], index in wordMap"> <!-- v-for state === unknown word, class = unknownword  -->
          <div :style="{color : activeColor}" :id="`container-${index}`">{{ word }}   <!--XXX possibly add a check in here to change the colour accordingly-->
          <linxreader :renderWord="word" @linxreaderload="LinxReaderLoaded"> </linxreader>
        </div>
        </div> <!-- TODO TODO replace wordMap with simply a list of words in the text file, then iterate through each word, make a call back to the server, have the server return a JSON object for each word. Before all of this create a default JSON object to initalize each word, each local JSON object should then be passed back to server for storage.-->
    </div>



  </template>
 //XXX Might be better to iterate over the map in the method beforehand
<script>
    import linximporter from "../components/Linx.vue"
    import linxreader from "../components/LinxReader.vue"
    import axios from 'axios';
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
          word : "",
          linxWord: ""
          }//one sentence class holds multiple sentenceitems
          //TODO keeping wordState in data return might not be necessary. Keep an array of both the word and the wordState in data return and simply modify the state within a method to allow for freezing. It should still update the rendering when one of the elements of the array changes.
          
        },
        components: {
        linximporter,
        linxreader
        },
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
        LinxReaderLoaded(linxWord) {
          this.linxWord = linxWord
        },
        computed: {
          RenderWordType() {
            var thing = document.getElementsByClassName('content');
            console.log(thing);
            return thing.word;
            console.log("asdALSDj");
          }
        }

        
    }
  
  }

</script>

<style>
.html {
  color: red;
}


</style>