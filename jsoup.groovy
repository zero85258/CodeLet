// @Grab('org.jsoup:jsoup:1.7.1')
// @Grab(group='org.apache.nifi', module='nifi-groovyx-processors', version='1.15.3')
@Grab(group='org.apache.nifi', module='nifi-api', version='1.15.3', scope='provided')
@Grab(group='org.jsoup', module='jsoup', version='1.11.3')

import groovy.json.JsonBuilder;
import org.jsoup.Jsoup
import org.jsoup.nodes.Document
import org.jsoup.nodes.Element
import org.jsoup.select.Elements
import org.apache.nifi.processor.io.StreamCallback
// System.exit(0)

final String WEB_URL = "http://.." //注释掉网站地址信息
// doc = Jsoup.connect(WEB_URL+"/html/cn/cn_newcenter/cn_newcenter_gsxw/index.html").get();
doc = Jsoup.connect("https://tw.yahoo.com/").get();

println(doc.title());
Elements eltes = doc.select("#nav > div > div:nth-child(1) > ul:nth-child(1)");
eltes.each {
    Element element -> println(element.text())
}

for (Element headline : eltes) {
    //获取新闻的二级页面
    println(headline.body().text());
    String url = headline.attr("href");
    //打开新闻二级页面
    Document doc1 =  Jsoup.connect(WEB_URL+url).get();
    //println("二级页面地址"+WEB_URL+url)
    Elements pageElets = doc1.select(".page > script");
    if (pageElets==null || pageElets.isEmpty()){
        System.exit(0)
    }

    def pageStr = pageElets.get(0).html();
    def match = "\\d,\\d";
    def results = (pageStr =~ /$match/)[0].toString();
    if (!results.find()){
        continue;
    }
    def pagenumber = results.split(",")
    for(int i=Integer.parseInt(pagenumber[1]);i<=Integer.parseInt(pagenumber[0]);i++){
        String pageHtml = 'index_';
        if(i!=1){
            pageHtml+=i
        }
        pageHtml+='.html'
        def pageUrl;
        if(i!=1){
            pageUrl = (WEB_URL+url).replace('index.html', pageHtml);
        }else{
            pageUrl = WEB_URL+url;
        }
        //读取每页新闻的连接 读取内容
        Document doc3 =  Jsoup.connect(pageUrl).get();
        Elements eles3 = doc3.select(".newribox >ul > li > a")
        eles3.each {Element e ->
            // def title = ;
            def contentUrl = WEB_URL+e.attr("href")
            // println('连接'+e.attr("href"))
            // println('标题'+e.text())
            //读取文章类容
            try {
                Document doc4 =  Jsoup.connect(contentUrl).get();
                def content = doc4.select(".newsmore").text();
                def date = doc4.select(".sumain > .newsdiv >p").get(0).text();
                println("文章内容："+content)
                // def map = [title:title,url:contentUrl,content:content,date:date]
                // def ruuid;

                def jsonstr = new JsonBuilder([
                    title: title = e.text(),
                    url: contentUrl,
                    id: ruuid  = UUID.randomUUID().toString(),
                    content: content,
                    date: date
                ]).toPrettyString()

                // def jsonstr = '{"title":"'+title+'",' +'"url":"'+contentUrl+'",'+'"id":"'+ruuid+'",'+'"content":"'+content+'",'+'"date":"'+date+'"}'
                flowFile = session.create();
                flowFile = session.write(flowFile, {inputStream, outputStream ->
                    // Select/project cols
                    outputStream.write(jsonstr.getBytes('UTF-8'));
                } as StreamCallback);
                session.transfer(flowFile, REL_SUCCESS)
            } catch (Exception exception){
                Thread.sleep(1)
            }
            // Thread.sleep(1000)
        }
    }
}
