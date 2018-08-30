import 'package:flutter/material.dart';
import 'package:dio/dio.dart';

String info;

void main(){
  checkIp("http://localhost:8080/api/ipapp/taobao?ip=192.34.23.23");
  runApp(new MyApp());
} 

class MyApp extends StatelessWidget{
  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return new MaterialApp(
      title: 'IpChcker',
      theme: new ThemeData(
        primarySwatch:Colors.orange,
      ),
      debugShowCheckedModeBanner: false,
      home: new MyHomePage(title:"IpChecker"),
    );
  }
}

class MyHomePage extends StatefulWidget{
  MyHomePage({Key key, this.title}):super(key:key);
  final String title;
  @override
  State<StatefulWidget> createState() {
    // TODO: implement createState
    return MyHomePageState();
  }

}
class MyHomePageState extends State<MyHomePage>{
  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return Scaffold(appBar: AppBar(
      title:Text(widget.title),
    ),
    body: Container(
      child:Center(
        child:Text(info),
      )
    ),);
  }
}

checkIp(String url) async{
  Dio dio = new Dio();
  Response response;
  try{
    response = await dio.request(url);
    info = response.data.toString();
    print(response.data.toString());
  }on DioError catch(e){
    e.toString();
  }
}