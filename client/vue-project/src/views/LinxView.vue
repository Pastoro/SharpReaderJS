<template>
    <div id="linx">
        
        <br>
      <linximporter v-if="!menuDisplay" @load="LoadHandler"></linximporter> <!-- TODO assign dynamic id -->
        <div class="grid-container"  > <!-- v-for state === unknown word, class = unknownword  -->

          <!-- <div :style="{color : activeColor}" :id="`container-${index}`">{{ word }}   XXX possibly add a check in here to change the colour accordingly-->
 
          <linxreader    
          v-for="(element,index) in words"
          :id="`linxreader${index}`"
          word="element.word"
          :renderWord="element.word" 
          :learned="element.state"
          :class="linxreadermethod(element.word,element.state)"
          @click="() => {displayIndex = index; element.state = ClickEvent(element.state);getDictionaryResult(element.word); menuDisplay = true}"
          >

          </linxreader>

        <Teleport v-if="menuDisplay" :to="`#linxreader${displayIndex}`">
          <LinxMenu
          v-if="menuDisplay"
          :queryResult="returnWord" 
          :searchedTerm="menuWord" 
          >
          </LinxMenu> 
        </Teleport>
        </div>
        





      </div>


  </template>
<script> //XXX you can probably use CSS to move the menu component around whilst only having it be bound to the grid-container class
    import linximporter from "../components/Linx.vue"
    import linxreader from "../components/LinxReader.vue"
    import axios from 'axios';
    import LinxMenu from "../components/LinxMenu.vue";
import { objectToString } from "@vue/shared";

//XXX Probably just have to make the sizing of the CSS grid elements dynamic
    export default{

        data () {

          return {
           // renderWord: '',
           // learned: '',
          returnWord : "",
          menuDisplay : false,
          displayIndex: 0,
          menuWord : "",
          classValue: "",
          activeColor: 'red',
          mywordState : null,
          rawhtml : '<span style="color: red">This should be red.</span>',
          wordMap : Map,
          separateState : {},
          words: [
            //arrayword:"", arraystate:{}
          ]
          }
          
        },
        components: {
    linximporter,
    linxreader,
    LinxMenu,
    LinxMenu,
    LinxMenu
}
        ,
        methods: {
          async getDictionaryResult(term){ //TODO supposedly computed properties are better and watch is discouraged
            //TODO Leave this as is for now, assuming we're gonna' use a third-party dictionary for this anyways.
            console.log(term);
            console.log("FIRED");

          var returnType = ((await axios({ method: "GET", params: {QueryWord: term}, responseType: 'json', url: " http://localhost:8030/fetchWord"})).data);
          console.log(typeof returnType);
          console.log(typeof term);
          this.menuWord = term;
          this.returnWord = returnType;
          //this.menuDisplay = !this.menuDisplay;
          //console.log(ctx.teleports);
          //return returnType;
          },
          LoadHandler(textMap, wordState) {
            //this.wordMap = textMap;
            this.mywordState = wordState;
            console.log(this.mywordState);
            var word;
            var state;

            for (const [key, value] of textMap) {
              word = key;
              state = value;
              this.words.push({word, state});
              console.log(this.words);
            }
            console.log(this.words);
            //this.words.forEach(word => {console.log(word)});
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
        linxreadermethod: function(renderWord, state) {
            var classA = '';
            console.log(renderWord);
            console.log(state);
            if (state === this.mywordState.DefaultState) {
              classA = 'wordthingUnknown';
            }
            else if (state === this.mywordState.LookedUpState) {
              classA = 'wordthingKnown';
            }
            return classA;
          },
          ClickEvent(state) {


            console.log(state);
            state = this.mywordState.LookedUpState;
            return state;
          }

      },
        computed: {

        },
        mounted(){
          console.log("MOUNTED HERE");
        }

        
    }
  
  

</script>

<style>
.grid-container {
  display: grid;
  grid-template-columns: repeat(10, 1fr);
  background-color: #10436d;
  padding: 10px;
  grid-auto-rows: repeat(1fr,120px 60px 120px);
  row-gap: 30px;
  column-gap: 30px;
}
.wordthingKnown {
  display:inline;
  background-color: rgba(255, 0, 34, 0.8);
  border: 1px solid rgba(25, 0, 255, 0.8);
  font-size: 30px;
  text-align: center;
  font: bold;
}
.wordthingUnknown {
  display:inline;
  background-color: rgba(0, 255, 179, 0.8);
  border: 1px solid rgba(0, 0, 0, 0.8);
  font-size: 30px;
  text-align: center;
}




</style>