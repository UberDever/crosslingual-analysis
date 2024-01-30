```html
<!-- index.html -->
<form>
    <div style="overflow:auto;"> 
        <div style="float:right;"> 
            <button type="button" id="prevBtn">Previous</button> 
        </div>
    </div>
</form>
```

```js
// script.js
function showTab(n) { 
  if (n == 0) { 
    document.getElementById("prevBtn").style.display = "none"; 
  } else { 
    document.getElementById("prevBtn").style.display = "inline"; 
  } 
} 
```