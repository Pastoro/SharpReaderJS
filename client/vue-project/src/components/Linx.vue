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
            reader.onload = e => {console.log(e.target.result); textContent = e.target.result; this.TextHandler(textContent)};
            //reader.onload = e => {console.log(e.target.result); textContent = e.target.result;textArray = new Array(this.textIntoArray(textContent).length); textArray = this.textIntoArray(textContent)};
            //reader.onloadend = () => {console.log(textArray[1])};
            //reader.addEventListener("loadend",() => console.log(textArray[1]));
            
            
            
        },
        //TODO check how to avoid scripts being run when uploading file from user
        //TODO also split string on newline currently it will add two words to one element if there is no space but there is an empty lin        
        TextHandler(textContent){
            const wordState = Object.freeze({
            DefaultState: 0,
            LookedUpState: 1,
            LearnedState: 2
          
            });
                //from each paragraph into each sentence into each word
            var paragraphArray = textContent.split(/\r?\n/);
            console.log(paragraphArray);


            const textMap = new Map();
            var sentenceArray = [];
            var textContentAltered = textContent.replace(/\./g,'');
            var wordArray = textContentAltered.split(" ");
            paragraphArray.forEach((paragraph, index) => {
                sentenceArray[index] = paragraph.split("."); //split removes the period
            });
            console.log(sentenceArray);
            console.log(wordArray);
            //TODO this should return sentences instead. Probably create sentence divs in LinxView instead.
            wordArray.forEach(word => {
                textMap.set(word, wordState.DefaultState);
                console.log(word);
            });
            this.$emit("load", textMap, wordState, wordArray);
            return textMap;
        }
    }
}


</script>