<template>
    <label class="linximporter">
        <input type="file" @change="LoadFile">
    </label>
</template>

<script>

export default{
    name:'linxUpload',
    data() {

        return{

    }
    },
    methods:{
        LoadFile(e){
            const file = e.target.files[0];
            const fileSize = file.size;
            var textContent = "";
            var textArray = [];
             
            if (file.size > 50000000)
            {
                console.error(`Maximum filesize exceeded: is ${fileSize}`);
                throw 0;
            }
            console.log(fileSize);
            const reader = new FileReader();
            reader.readAsText(file);
            reader.onload = e => {console.log(e.target.result); textContent = e.target.result; this.TextIntoMap(textContent)};
            //reader.onload = e => {console.log(e.target.result); textContent = e.target.result;textArray = new Array(this.textIntoArray(textContent).length); textArray = this.textIntoArray(textContent)};
            //reader.onloadend = () => {console.log(textArray[1])};
            //reader.addEventListener("loadend",() => console.log(textArray[1]));
            
            
            
        },
        //TODO check how to avoid scripts being run when uploading file from user
        //TODO also split string on newline currently it will add two words to one element if there is no space but there is an empty lin        
        TextIntoMap(textContent){
            const sentence = {
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
            const wordState = Object.freeze({
            DefaultState: 0,
            LookedUpState: 1,
            LearnedState: 2
          
            });
            const textMap = new Map();
            var wordArray = textContent.split(" ");
                //TODO reimplement stemmer before this is done
            //TODO this should return sentences instead
            wordArray.forEach(word => {
                textMap.set(word, wordState.DefaultState);
            });
            this.$emit("load", textMap, wordState);
            return textMap;
        }
    }
}


</script>