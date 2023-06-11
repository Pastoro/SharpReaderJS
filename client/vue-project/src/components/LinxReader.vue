<template>
<label class = "linxreader">

<button 
@click="isActive=!isActive"
>

{{ renderWord }} </button>
<LinxMenu v-if="isActive" :searchedTerm="renderWord" :queryResult="returnedWord" ></LinxMenu>
</label>

</template>



<script>

    import LinxMenu from './LinxMenu.vue';

    export default {
      components: {
    LinxMenu
},
        props: {renderWord : String,
                wordId : String,
                learned : Number,
                created() {
                    this.wordReturn = this.renderWord;
                    this.idReturn = this.wordId;
                    this.learnedReturn = this.learned;
                }
            },
        data() {
            return {
                isActive: false,
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




        }
        
    }


</script>

<style>


</style>