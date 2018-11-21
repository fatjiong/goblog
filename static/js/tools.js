// 给文章点赞/踩
function addArticleCounter(id,source) {
    $.post("/article/counter",{id:id,source:source},function(res){
        location.reload();
    },"json");
}