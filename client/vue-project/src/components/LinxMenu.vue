
<template>
    <label class ="linxmenu">



        <div class="linxmenuPopup" >

            <p > {{ searchedTerm}} </p>

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
        this.getDictionaryResult(this.searchedTerm);
    },
    methods:{
        getDictionaryResult(term){ //TODO supposedly computed properties are better and watch is discouraged
            console.log(term);
            console.log("FIRED");
            var returnType = "";
            axios({ method: "GET", params: {QueryWord: term}, responseType: 'json', url: " http://localhost:8030/fetchWord"}).then(function (response) {
            console.log(response.data); returnType = response.data;
            //console.log(JSON.stringify(JSON.parse(response), null, 2));
          });
          this.$emit('gotDictionaryResult',returnType);
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


</style>