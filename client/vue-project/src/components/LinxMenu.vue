
<template>
    <label class ="linxmenu">



        <div class="linxmenuPopup" >

            <p class="search"> {{ searchedTerm }} </p>
            <p class="query"> {{ queryResult }}</p>

        </div>



        
    </label>




</template>

<script>
import axios from 'axios';


export default{
    
    data(){
        return{
            searchTerm: "",
            resultTerm: ""
        }
    },
    props: {
        searchedTerm: String,
        queryResult: String,
        
        created(){
            this.searchTerm = this.searchedTerm;
            this.resultTerm = this.queryResult;
        }
    },
    mounted() {
        this.resultTerm = this.getDictionaryResult(this.searchedTerm);
    },
    methods:{
        async getDictionaryResult(term){ //TODO supposedly computed properties are better and watch is discouraged
            //TODO Leave this as is for now, assuming we're gonna' use a third-party dictionary for this anyways.
            console.log(term);
            console.log("FIRED");

            var returnType = ((await axios({ method: "GET", params: {QueryWord: term}, responseType: 'json', url: " http://localhost:8030/fetchWord"})).data);
          console.log(returnType);
          this.$emit("gotDictionaryResult",returnType);
          return returnType;
        }
    }
}



</script>



<style>

.linxmenuPopup {
    color: blueviolet;
    text-shadow: 18cm;
    background-color: whitesmoke;
    width: 400px;
    height: 500px;
    z-index: 30;
    position: absolute;
    
}
.query{
    color:maroon;
}


</style>