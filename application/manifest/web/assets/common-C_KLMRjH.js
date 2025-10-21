const g=(t,s=1)=>{t=t.replace(/^#/,""),t.length===3&&(t=t.split("").map(o=>o+o).join(""));const n=parseInt(t,16),r=n>>16&255,a=n>>8&255,c=n&255;return`rgba(${r}, ${a}, ${c}, ${s})`};export{g as h};
