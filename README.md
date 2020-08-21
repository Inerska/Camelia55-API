
<center>  
 <a href="https://github.com/Inerska/Camelia55-API">  
 <img src="https://raw.githubusercontent.com/Inerska/Camelia55-API/master/cam55logo.png" alt="Logo">  
 </a>  
 <h4>Camélia55-API is no way affiliated with Camélia55, please not consider it as its official API</h4>  
 <hr />
 <h5>Camélia55-API is free and open-source, do not hesitate to contribute, fork or star it.<br/> You can use any programming language you want.</h5>
 <p>The <b>A</b>pplication <b>P</b>rogramming <b>I</b>nterface is made for educational purpose. It's aiming to be useful and easy to use. The repository contains a ready-to-use web server, host it and use it. <br/><i>n.b. The delay between requests is set to 3 seconds</i> </p>
</center>
## Use
<pre>
$ mkdir Camel55
$ cd Camel55
$ git clone https://github.com/Inerska/Camelia55-API.git
$ sudo apt install golang-go
$ go run main.go
</pre>
<br>
- Retrieving staff information
<pre>GET http://127.0.0.1:8080/api/v1/users/staff</pre>
<i>n.b. Use 127.0.0.1 only if you are hosting the web server locally</i>
