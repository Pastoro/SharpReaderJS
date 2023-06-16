<template>
<label class = "linxreader">

<span 
@click="onClick()"
>

{{ renderWord }} </span>

</label>

</template>



<script>
//TODO I have to either rewrite my component somehow. or to make the dynamic positioning funciton I would have to rewrite that portion of it. Maybe you can set a separate active object with which to position your popup with.
    import { vModelCheckbox } from 'vue';
import LinxMenu from './LinxMenu.vue';
    var renderedMenusIndex = 0;
    var renderedMenusArray = [];


    export default {
      components: {
    LinxMenu
},
        props: {renderWord : String,
                wordId : String,
                learned : Number,
                isMounted : Boolean,
                created() {

                    this.wordReturn = this.renderWord;
                    this.idReturn = this.wordId;
                    this.learnedReturn = this.learned;
                }
            },
        data() {
            return {
                index: 0,
                renderedMenus: [],
                returnWord: "",
                isActive: false,
                isActive2: false,
                wordReturn: "",
                idReturn: null,
                learnedReturn: 0,
                sentence : {
            sentenceitem : {
              unknownword : {
                //Colors, rendering stuff
              },
              knownword : {
                tha : String
                //Colors, rendering stuff
              }
            }
          }
            }
        },
        methods: {
          

          onClick(){
            if (!this.isActive && !this.isActive2) {
              this.isActive = true;
            }
            else if (this.isActive && !this.isActive2) {
              this.isActive = false;
              this.isActive2 = true;
            }
            else if (!this.isActive && this.isActive2) {
              this.isActive = true;
              this.isActive2 = false;
            }
            else if (this.isActive && this.isActive2) {
              this.isActive = false;
              this.isActive2 = false;
            }
            return this.index;
          }



        }
        
    }


</script>

<style>

.linxreader {
  position: relative;
}

</style>