package bonemaster

import (
	"html/template"
	"net/http"
	"time"
 //   "fmt"

	"appengine"
	"appengine/datastore"
)

// [START comment_struct]
type Comments struct {
	GivenName string
	EMail     string
	Comment   string
	Date      string
}

// [START comment_struct]
type Contacts struct {
     GivenName string
     EMail     string
     Comment   string
     Date      string
}

// [END greeting_struct]

func init() {
	http.HandleFunc("/", static)
	http.HandleFunc("/comments.html", root)
    http.HandleFunc("/contact.html", contact)
	http.HandleFunc("/sign", sign)
    http.HandleFunc("/postContactForm", postContactForm)
}

func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "files/"+r.URL.Path)
}

// commentKey returns the key used for all commentbook entries.
func commentKey(c appengine.Context) *datastore.Key {
	// The string "default_comment" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Comment", "default_comment", 0, nil)
}

// contactKey returns the key used for all commentbook entries.
func contactKey(c appengine.Context) *datastore.Key {
     // The string "default_comment" here could be varied to have multiple guestbooks.
     return datastore.NewKey(c, "Contact", "default_contact", 0, nil)
}

// [START func_root]
func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	// Ancestor queries, as shown here, are strongly consistent with the High
	// Replication Datastore. Queries that span entity groups are eventually
	// consistent. If we omitted the .Ancestor from this query there would be
	// a slight chance that Greeting that had just been written would not
	// show up in a query.
	// [START query]
	q := datastore.NewQuery("Comments").Ancestor(commentKey(c)).Order("-Date").Limit(5)
	// [END query]
	// [START getall]
	comments := make([]Comments, 0, 10)
	if _, err := q.GetAll(c, &comments); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// [END getall]
	if err := commentTemplate.Execute(w, comments); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// [START func_root]
func contact(w http.ResponseWriter, r *http.Request) {
//     c := appengine.NewContext(r)
     // Ancestor queries, as shown here, are strongly consistent with the High
     // Replication Datastore. Queries that span entity groups are eventually
     // consistent. If we omitted the .Ancestor from this query there would be
     // a slight chance that Greeting that had just been written would not
     // show up in a query.

     contacts := make([]Contacts, 0, 10)     
     
     if err := contactTemplate.Execute(w, contacts); err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
     }
}

// [END func_root]

var commentTemplate = template.Must(template.New("book").Parse(`
<!DOCTYPE html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>BoneMaster BBQ</title>
        <!-- Bootstrap -->
        <link href="css/bootstrap.min.css" rel="stylesheet">
        <link href="css/bootstrap-theme.min.css" rel="stylesheet">        
        <link href="css/bonemaster.css" rel="stylesheet">
        <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
        <link href="css/ie10-viewport-bug-workaround.css" rel="stylesheet" type="text/css"/>
        <link rel="icon" href="favicon1.png" sizes="16x16" type="image/png">
       <!-- Google Analytics -->
        <script>
        (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
        (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
        m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
        })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
        ga('create', 'UA-88295925-1', 'auto');  // Replace with your property ID.
        ga('send', 'pageview');
        </script>
        <!-- End Google Analytics -->
    </head>
    <body id="wrapper">
        <div class="container">
            <img src="images/BonemasterBBQ.jpg" class="img-rounded img-responsive center-block" alt="BoneMaster BBQ"/>
            <nav class="navbar navbar-default">
                <div class="container-fluid">
                    <!-- Brand and toggle get grouped for better mobile display -->
                    <div class="navbar-header">
                        <button type="button" class="btn-default navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
                            <span class="sr-only">Toggle navigation</span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                        </button>                        
                    </div>
                    <!-- Collect the nav links, forms, and other content for toggling -->
                    <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                        <ul class="nav navbar-nav">                            
                            <li><a href="index.html"><h4>Home</h4></a></li>
                            <li><a href="about.html"><h4>About</h4></a></li>
                            <li><a href="menu.html"><h4>Menu</h4></a></li>
                            <li><a href="photos.html"><h4>Photos</h4></a></li>
                            <li><a href="contact.html"><h4>Contact</h4></a></li>
                            <li class="active"><a href="comments.html"><h4>Comments</h4></a></li>
                            <li><a href="catering.html"><h4>Catering</h4></a></li>
                            <li>
                                <a href="#modalSearch" data-toggle="modal" data-target="#modalSearch"><span id="searchGlyph" class="glyphicon glyphicon-search"></span> <span class="hidden-sm">Search</span></a>
                            </li>                            
                        </ul>
                    </div><!-- /.navbar-collapse -->
                </div><!-- /.container-fluid -->
            </nav>
            <!-- End Nav Content -->
            <!-- Begin Modal Search -->
            <div id="modalSearch" class="modal fade" role="dialog">
                <div class="modal-dialog">
                    <!-- Modal content-->
                    <div class="modal-content">
                        <div class="modal-header">
                            <button type="button" class="close" data-dismiss="modal">&times;</button>
                            <h4 class="modal-title">Search BonemasterBBQ</h4>
                        </div>
                        <div class="modal-body">
                            <script>                                
                                (function() {
                                    var cx = '006257532291329875067:0tnlhh1agjw';
                                    var gcse = document.createElement('script');
                                    gcse.type = 'text/javascript';
                                    gcse.async = true;
                                    gcse.src = 'https://cse.google.com/cse.js?cx=' + cx;
                                    var s = document.getElementsByTagName('script')[0];
                                    s.parentNode.insertBefore(gcse, s);
                                })();
                            </script>
                            <gcse:search></gcse:search>
                            <!-- These styles fix CSE and Bootstrap 3 conflict -->
                            <style type="text/css">
                                .reset-box-sizing, .reset-box-sizing *, .reset-box-sizing *:before, .reset-box-sizing *:after,  .gsc-inline-block
                                {
                                    -webkit-box-sizing: content-box;
                                    -moz-box-sizing: content-box;
                                    box-sizing: content-box;
                                }
                                input.gsc-input, .gsc-input-box, .gsc-input-box-hover, .gsc-input-box-focus, .gsc-search-button
                                {
                                    box-sizing: content-box;
                                    line-height: normal;
                                }
                            </style>
                        </div>
                        <div class="modal-footer">
                            <button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
                        </div>
                    </div>
                </div>
            </div>
            <!--// End Search Modal -->
            <!-- Begin Main Content -->                           
            <div class="container" id="maincontent">
                <div class="jumbotron">                  
                    <h1 class="text-center">Comments</h1>
                </div>
                <hr>
                <!-- Begin Database Input -->
                <div>
                    {{range .}}                            
                        <h2>{{.GivenName}}</h2>                                       
                        <h3>{{.Comment}}</h3>
                        <h4>{{.Date}}</h4>
                        <hr />
                    {{end}}
                </div>
                <!-- //End Database Input -->
                <form action="/sign" method="post" class="form-horizontal">
                    <div class="form-group">
                        <label for="givenname">First and Last Name:</label>
                        <input type="text" class="form-control"  placeholder="First and Last Name" name="givenname"/>
                    </div>
                    <div class="form-group">
                        <label for="email">Email address:</label>
                        <input type="text" class="form-control" placeholder="Email" name="email" />
                    </div>          
                    <label for="comment">Comment:</label>
                    <textarea class="form-control" placeholder="Comment" name="commentFormcontent" rows="5"></textarea>
                    <button type="submit" class="btn btn-primary" value="Submit">Submit</button>
                </form>                
                <hr>
                <!-- Start Footer -->   
                <div id="footer">
                    <h4>&copy; Copyright 2012 <a href="mailto:ron@bonemasterbbq.com"><u>ron@bonemasterbbq.com</u>.</a>  Designed by Jason Bauer and Brett Salemink.</h4>
                </div>
                <!-- End Footer --> 
            </div>
            <!-- End Main Content -->
        </div>
        <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
        <!-- Include all compiled plugins (below), or include individual files as needed -->
        <script src="js/bootstrap.min.js"></script>    
    </body>
</html>
`))

var contactTemplate = template.Must(template.New("contact").Parse(`
<!DOCTYPE html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>BoneMaster BBQ</title>
        <!-- Bootstrap -->
        <link href="css/bootstrap.min.css" rel="stylesheet">
        <link href="css/bootstrap-theme.min.css" rel="stylesheet">        
        <link href="css/bonemaster.css" rel="stylesheet">
        <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
        <link href="css/ie10-viewport-bug-workaround.css" rel="stylesheet" type="text/css"/>
        <link rel="icon" href="favicon1.png" sizes="16x16" type="image/png">
        <!-- Google Analytics -->
        <script>
        (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
        (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
        m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
        })(window,document,'script','//www.google-analytics.com/analytics.js','ga');
        ga('create', 'UA-88295925-1', 'auto');  // Replace with your property ID.
        ga('send', 'pageview');
        </script>
        <!-- End Google Analytics -->
    </head>
    <body id="wrapper">
        <div class="container">
            <img src="images/BonemasterBBQ.jpg" class="img-rounded img-responsive center-block" alt="BoneMaster BBQ"/>
            <!-- Begin Nav Content -->
            <nav class="navbar navbar-default">
                    <div class="container-fluid">
                        <!-- Brand and toggle get grouped for better mobile display -->
                        <div class="navbar-header">
                            <button type="button" class="btn-default navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
                                <span class="sr-only">Toggle navigation</span>
                                <span class="icon-bar"></span>
                                <span class="icon-bar"></span>
                                <span class="icon-bar"></span>
                            </button>                        
                        </div>
                        <!-- Collect the nav links, forms, and other content for toggling -->
                        <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                            <ul class="nav navbar-nav">                            
                                <li><a href="index.html"><h4>Home</h4></a></li>
                                <li><a href="about.html"><h4>About</h4></a></li>
                                <li><a href="menu.html"><h4>Menu</h4></a></li>
                                <li><a href="photos.html"><h4>Photos</h4></a></li>
                                <li class="active"><a href="contact.html"><h4>Contact</h4></a></li>
                                <li><a href="comments.html"><h4>Comments</h4></a></li>
                                <li><a href="catering.html"><h4>Catering</h4></a></li>
                                <li>
                                    <a href="#modalSearch" data-toggle="modal" data-target="#modalSearch"><span id="searchGlyph" class="glyphicon glyphicon-search"></span> <span class="hidden-sm">Search</span></a>
                                </li>                                                        
                            </ul>
                        </div><!-- /.navbar-collapse -->
                    </div><!-- /.container-fluid -->
                </nav>
                <!-- End Nav Content -->
                <!-- Begin Modal Search -->
                <div id="modalSearch" class="modal fade" role="dialog">
                    <div class="modal-dialog">
                        <!-- Modal content-->
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal">&times;</button>
                                <h4 class="modal-title">Search BonemasterBBQ</h4>
                            </div>
                            <div class="modal-body">
                                <script>                                
                                    (function() {
                                        var cx = '006257532291329875067:0tnlhh1agjw';
                                        var gcse = document.createElement('script');
                                        gcse.type = 'text/javascript';
                                        gcse.async = true;
                                        gcse.src = 'https://cse.google.com/cse.js?cx=' + cx;
                                        var s = document.getElementsByTagName('script')[0];
                                        s.parentNode.insertBefore(gcse, s);
                                    })();
                                </script>
                                <gcse:search></gcse:search>
                                <!-- These styles fix CSE and Bootstrap 3 conflict -->
                                <style type="text/css">
                                    .reset-box-sizing, .reset-box-sizing *, .reset-box-sizing *:before, .reset-box-sizing *:after,  .gsc-inline-block
                                    {
                                        -webkit-box-sizing: content-box;
                                        -moz-box-sizing: content-box;
                                        box-sizing: content-box;
                                    }
                                    input.gsc-input, .gsc-input-box, .gsc-input-box-hover, .gsc-input-box-focus, .gsc-search-button
                                    {
                                        box-sizing: content-box;
                                        line-height: normal;
                                    }
                                </style>
                            </div>
                            <div class="modal-footer">
                                <button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
                            </div>
                        </div>
                    </div>
                </div>
                <!--// End Search Modal -->
            <!-- Start Main Content -->                           
            <div class="container" id="maincontent">
                <div class="jumbotron">                    
                    <h1 class="text-center">Contact</h1>
                </div>
                <hr>          
                <form action="/postContactForm" method="post" class="form-horizontal">
                    <div class="form-group">
                        <label for="givenname">First and Last Name:</label>
                        <input type="text" class="form-control"  placeholder="First and Last Name" name="givenname"/>
                    </div>
                    <div class="form-group">
                        <label for="email">Email address:</label>
                        <input type="text" class="form-control" placeholder="Email" name="email" />
                    </div>              
                    <label for="comment">Comment/Orders:</label>
                    <textarea class="form-control" placeholder="Comments/Orders" name="contactFormcontent" rows="30"></textarea>
                    <button type="submit" class="btn btn-primary" value="Submit">Submit</button>
                </form>        
                <hr>
                <!-- Start Footer -->   
                <div id="footer">
                    <h4>&copy; Copyright 2012 <a href="mailto:ron@bonemasterbbq.com"><u>ron@bonemasterbbq.com</u>.</a>  Designed by Jason Bauer and Brett Salemink.</h4>
                </div>
                <!-- End Footer --> 
            </div>
            <!-- End Main Content -->
        </div>
        <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
        <!-- Include all compiled plugins (below), or include individual files as needed -->
        <script src="js/bootstrap.min.js"></script>
    </body>
</html>
`))

// [START func_sign]
func sign(w http.ResponseWriter, r *http.Request) {
	// [START new_context]
    c := appengine.NewContext(r)
	// [END new_context]

    // get current timestamp
    currentTime := time.Now().Local()
    currentDate := currentTime.Format("Mon Jan 2 15:04:05 MST 2006")


	g := Comments{
		GivenName: r.FormValue("givenname"),
		EMail:     r.FormValue("email"),
		Comment:   r.FormValue("commentFormcontent"),
		Date:      currentDate,
	}

	// We set the same parent key on every Greeting entity to ensure each Greeting
	// is in the same entity group. Queries across the single entity group
	// will be consistent. However, the write rate to a single entity group
	// should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "Comments", commentKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/comments.html", http.StatusFound)
	// [END if_user]
}

// [START func_sign]
func postContactForm(w http.ResponseWriter, r *http.Request) {
     // [START new_context]
    c := appengine.NewContext(r)
     // [END new_context]

    // get current timestamp
    currentTime := time.Now().Local()
    currentDate := currentTime.Format("Mon Jan 2 15:04:05 MST 2006")


     g := Comments{
          GivenName: r.FormValue("givenname"),
          EMail:     r.FormValue("email"),
          Comment:   r.FormValue("contactFormcontent"),
          Date:      currentDate,
     }

     // We set the same parent key on every Greeting entity to ensure each Greeting
     // is in the same entity group. Queries across the single entity group
     // will be consistent. However, the write rate to a single entity group
     // should be limited to ~1/second.
     key := datastore.NewIncompleteKey(c, "Contacts", contactKey(c))
     _, err := datastore.Put(c, key, &g)
     if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
     }
     http.Redirect(w, r, "/successcontactsubmission.html", http.StatusFound)
     // [END if_user]
}
 