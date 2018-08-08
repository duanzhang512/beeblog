{{define "navbar"}}
<nav class="navbar navbar-expand-lg navbar-light" style="background-color: #e3f2fd;">  
  <a clas s="navbar-brand" href="/">我的博客</a>
    <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
      <div class="navbar-nav">
      	<li {{if .IsHome}}class="active"{{end}}>
        	<a class="nav-item nav-link" href="/">首页 <span class="sr-only">(current)</span></a>
    	</li>
    	<li {{if .IsCategory}} class="active"{{end}}>
        	<a class="nav-item nav-link" href="/category">分类</a>
        </li>
        <li {{if .IsTopic}} class="active"{{end}}>
        	<a class="nav-item nav-link" href="/topic">文章</a>
        </li>
      </div>
    </div>
    
    <div class="pull-right">
		<ul class="nav navbar-nav">
			{{if .IsLogin}}
				<li><a href="/login?exit=true">退出</a></li>
			{{else}}
				<li><a href="/login">管理员登录</a></li>
			{{end}}
		</ul>
	</div>
</nav>
{{end}}    