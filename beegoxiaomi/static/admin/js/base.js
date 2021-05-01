$(function(){	
	app.init()
	
	$(window).rezize(function(){
		app.resizeIframe()
	})
	
})
var config={
	adminPath:"beego_admin"
}

var app={
	init:function(){
		this.slideToggle();
		this.resizeIframe();
		this.confirmDelete();
		this.changeStatus();
		this.changeNum();
	},	
	slideToggle:function(){
		$('.aside h4').click(function(){
		
			$(this).siblings('ul').slideToggle();
		})
	},
	resizeIframe:function(){		
		$("#rightMain").height($(window).height()-80)
	},
	confirmDelete:function(){
		$(".delete").click(function(){
			var flag = confirm("您确定要删除吗")
			return flag
		})
	},
	changeStatus:function(){
		$(".chStatus").click(function(){
			var id = $(this).attr("data-id");
			var table = $(this).attr("data-table");
			var field = $(this).attr("data-field");
			var el =$(this);
			$.get("/"+config.adminPath+"/main/changeStatus",{id:id,table:table,field:field},function(response){
				console.log(response);
				if(response.success){
					if(el.attr("src").indexOf("yes") != -1){
						el.attr("src","/static/admin/images/no.gif")
					}else{
						el.attr("src","/static/admin/images/yes.gif")
					}
				}
			})
		})
	},
	changeNum:function(){
		/*
		1、获取el里面的值  var spanNum=$(this).html()


		2、创建一个input的dom节点   var input=$("<input value='' />");


		3、把input放在el里面   $(this).html(input);


		4、让input获取焦点  给input赋值    $(input).trigger('focus').val(val);

		5、点击input的时候阻止冒泡 

					$(input).click(function(e){
						e.stopPropagation();				
					})

		6、鼠标离开的时候给span赋值,并触发ajax请求

			$(input).blur(function(){
				var inputNum=$(this).val();
				spanEl.html(inputNum);
				触发ajax请求
				
			})
		*/

		$(".chSpanNum").click(function(){
			var id=$(this).attr("data-id");
			var table=$(this).attr("data-table");
			var field=$(this).attr("data-field");
			var spanNum=$(this).attr("data-num");
			var spanEl=$(this)  //保存span这个dom节点

			var input=$("<input value='' style='width:60px' />");
			$(this).html(input);
			$(input).trigger('focus').val(spanNum);   //让输入框获取焦点并设置值
			$(input).click(function(e){
				e.stopPropagation();				
			})
			$(input).blur(function(e){
				var inputNum=$(this).val();
				spanEl.html(inputNum);
				//异步请求修改数量
				$.get("/"+config.adminPath+"/main/editNum",{id:id,table:table,field:field,num:inputNum},function(response){				
					if(!response.success){
						console.log(response)
					}
				})
			})
		})
	}
}
