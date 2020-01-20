# Create your views here.
from django.http import HttpResponse


def index(request):
    return HttpResponse("Hello, world. You're at the polls index.")









<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">

<html>
<head>
<title>
삼성전자/분기보고서/2019.11.14
</title>
<meta content="IE=EmulateIE8" http-equiv="X-UA-Compatible"/>
<meta content="text/html; charset=utf-8" http-equiv="Content-Type"/>
<style type="text/css">
    .x-window-dlg .ext-mb-download {
        background:transparent url(images/download.gif) no-repeat top left;
        height:46px;
    }
</style>
<link charset="utf-8" href="/css/viewer.css" rel="stylesheet" type="text/css"/>
<script src="/js/prototype.js" type="text/javascript"></script>
<script src="/js/jquery/jquery-all.js" type="text/javascript"></script>
<!-- 2014.03.11 ext 3.4 -->
<!--[if lte IE 8]><link rel="stylesheet" type="text/css" href="/js/ext-viewer/resources/css/ext-all-ie8.css" ><![endif]-->
<!--[if (gte IE 9)|!(IE)]><!--><link href="/js/ext-viewer/resources/css/ext-all.css" rel="stylesheet" type="text/css"/><!--<![endif]-->
<script src="/js/ext-viewer/adapter/ext/ext-base.js" type="text/javascript"></script>
<script src="/js/ext-viewer/ext-all.js" type="text/javascript"></script>
<!-- application css -->
<link href="/css/report_c.css" media="screen" rel="stylesheet" type="text/css"/>
<link charset="utf-8" href="/css/viewer/report_dart.css" rel="stylesheet" type="text/css"/>
<link charset="utf-8" href="/css/viewer/report_pop.css" rel="stylesheet" type="text/css"/>
<!-- x-xeries js libraries  -->
<script src="/js/xjs.js" type="text/javascript"></script>
<script src="/js/dart/build/ds-all-min.js" type="text/javascript"></script>
<script src="/js/viewer/popup.js" type="text/javascript"></script>
<script src="/js/idart/idart.js" type="text/javascript"></script><!-- idart용 -->
<script type="text/javascript">

var winTestNotice = null;	// xwindow object
var divTestNotice = "divTestNoticeWin";	// ajax 수신 div
$j(document).ready(function() {
	initPage();

	initScroll();
	$j(window).resize(function(){
		$j("#center-panel > div > div").resize(resizeFrame);
		resizeFrame();
	}).resize();
});


var resizeFrame = function() {
	w	= parseInt($j("#center-panel > div > div").css("width"));
	h	= parseInt($j("#center-panel > div > div").css("height"));
	//$j("#ifrm").css("width", (w - 12) + "px");
	
		
	
	$j("#ifrm").css("height", h + "px");
	$j("#mobileScroll").css("height", h + "px");
};

var west;

//모바일(ios)에서 본문내용 스크롤 추가
function initScroll() {
	var iosKeywords = new Array("iPhone","iPod","iPad");
	for(var i=0; i<iosKeywords.length; i++){
		if(navigator.userAgent.indexOf(iosKeywords[i]) > -1) {
			$j("#mobileScroll").css("overflow","auto");
			$j("#mobileScroll").css("-webkit-overflow-scrolling","touch");
			$j("#mobileScroll").css("width","100%");
		}
	}
}

function initPage() {

	initCorpInfo();	// 기업개황팝업 초기화

	Ext.BLANK_IMAGE_URL = "/js/extjs/resources/images/default/s.gif";

	Ext.state.Manager.setProvider(new Ext.state.CookieProvider());

	var north = new Ext.BoxComponent({
	    region:'north',
	    el: 'north',
	    height:74
	});

	var Tree = Ext.tree;

	var treeRoot = new Tree.TreeNode({
		text: "전체",
		id: "root",
		href: "javascript: viewDoc('20191114001273', '6958001', null, null, null, 'dart3.xsd')"
	});

	west = new Tree.TreePanel({
		id:"west-panel",
		title:"문서목차",
		region:"west",
		split:true,
		autoScroll:true,
		width: 200,
		minSize: 175,
		maxSize: 400,
		collapsible: true,
		collapsed: false,
		margins:"5 0 5 5",
		animate:false,
		useArrows:false,
		loader: new Tree.TreeLoader(),
		enableDD:false,
		containerScroll: true,
		checked: false,
		rootVisible: false,
		titleCollapse: false
	});

	var cnt = 0;

	
	
		// 1
		treeNode1 = new Tree.TreeNode({
			text: "분 기 보 고 서",
			id: "1",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '1', '670', '4055', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeRoot.appendChild(treeNode1);
		
			
			
		
	
		// 2
		treeNode1 = new Tree.TreeNode({
			text: "【 대표이사 등의 확인 】",
			id: "2",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '2', '12757', '401', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeRoot.appendChild(treeNode1);
		
	
		// 3
		treeNode1 = new Tree.TreeNode({
			text: "I. 회사의 개요",
			id: "3",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '3', '13162', '309931', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeNode2 = new Tree.TreeNode({
			text: "1. 회사의 개요",
			id: "4",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '4', '13252', '209842', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "2. 회사의 연혁",
			id: "5",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '5', '223098', '17213', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "3. 자본금 변동사항",
			id: "6",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '6', '240315', '4054', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "4. 주식의 총수 등",
			id: "7",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '7', '244373', '61513', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "5. 의결권 현황",
			id: "8",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '8', '305890', '7341', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "6. 배당에 관한 사항 등",
			id: "9",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '9', '313235', '9844', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeRoot.appendChild(treeNode1);
		
	
		// 4
		treeNode1 = new Tree.TreeNode({
			text: "II. 사업의 내용",
			id: "10",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '10', '323097', '276444', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeRoot.appendChild(treeNode1);
		
	
		// 5
		treeNode1 = new Tree.TreeNode({
			text: "III. 재무에 관한 사항",
			id: "11",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '11', '599545', '1073012', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeNode2 = new Tree.TreeNode({
			text: "1. 요약재무정보",
			id: "12",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '12', '599672', '21551', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "2. 연결재무제표",
			id: "13",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '13', '621477', '85372', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "3. 연결재무제표 주석",
			id: "14",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '14', '706853', '392470', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "4. 재무제표",
			id: "15",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '15', '1099327', '69959', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "5. 재무제표 주석",
			id: "16",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '16', '1169290', '288475', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "6. 기타 재무에 관한 사항",
			id: "17",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '17', '1457795', '214748', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeRoot.appendChild(treeNode1);
		
	
		// 6
		treeNode1 = new Tree.TreeNode({
			text: "IV. 이사의 경영진단 및 분석의견",
			id: "18",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '18', '1672561', '269', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeRoot.appendChild(treeNode1);
		
	
		// 7
		treeNode1 = new Tree.TreeNode({
			text: "V. 감사인의 감사의견 등",
			id: "19",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '19', '1672834', '15154', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeRoot.appendChild(treeNode1);
		
	
		// 8
		treeNode1 = new Tree.TreeNode({
			text: "VI. 이사회 등 회사의 기관에 관한 사항",
			id: "20",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '20', '1687992', '116730', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeNode2 = new Tree.TreeNode({
			text: "1. 이사회에 관한 사항",
			id: "21",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '21', '1688105', '82574', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "2. 감사제도에 관한 사항",
			id: "22",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '22', '1770683', '33200', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "3. 주주의 의결권 행사에 관한 사항",
			id: "23",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '23', '1803887', '821', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeRoot.appendChild(treeNode1);
		
	
		// 9
		treeNode1 = new Tree.TreeNode({
			text: "VII. 주주에 관한 사항",
			id: "24",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '24', '1804726', '56670', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeRoot.appendChild(treeNode1);
		
	
		// 10
		treeNode1 = new Tree.TreeNode({
			text: "VIII. 임원 및 직원 등에 관한 사항",
			id: "25",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '25', '1861400', '805938', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeNode2 = new Tree.TreeNode({
			text: "1. 임원 및 직원의 현황",
			id: "26",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '26', '1861512', '791708', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "2. 임원의 보수 등",
			id: "27",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '27', '2653224', '14100', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeRoot.appendChild(treeNode1);
		
	
		// 11
		treeNode1 = new Tree.TreeNode({
			text: "IX. 계열회사 등에 관한 사항",
			id: "28",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '28', '2667342', '504897', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeRoot.appendChild(treeNode1);
		
	
		// 12
		treeNode1 = new Tree.TreeNode({
			text: "X. 이해관계자와의 거래내용",
			id: "29",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '29', '3172243', '56906', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeRoot.appendChild(treeNode1);
		
	
		// 13
		treeNode1 = new Tree.TreeNode({
			text: "XI. 그 밖에 투자자 보호를 위하여 필요한 사항",
			id: "30",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '30', '3229153', '80738', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeRoot.appendChild(treeNode1);
		
	
		// 14
		treeNode1 = new Tree.TreeNode({
			text: "【 전문가의 확인 】",
			id: "31",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '31', '3309895', '346', 'dart3.xsd');}
			}
		});
		cnt++;
		
		treeNode2 = new Tree.TreeNode({
			text: "1. 전문가의 확인",
			id: "32",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '32', '3309996', '112', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeNode2 = new Tree.TreeNode({
			text: "2. 전문가와의 이해관계",
			id: "33",
			cls: "text",
			listeners: {
				click: function() {viewDoc('20191114001273', '6958001', '33', '3310112', '115', 'dart3.xsd');}
			}
		});
		cnt++;
			
		treeNode1.appendChild(treeNode2);
		
		treeRoot.appendChild(treeNode1);
		
	

	west.setRootNode(treeRoot);

	var viewport = new Ext.Viewport({
	     layout:'border',
	     items:[
	     	north,
	     	west,
	     	new Ext.Panel({
	             region:'center',
	             id:'center-panel',
	             margins:'5 5 5 5',
	             autoScroll:false,
	             items: [{
	                 contentEl: 'center',
	                 border:false,
	                 autoScroll:false
	             }]
	         })
	      ]
	 });



	if (cnt > 200) {
		treeRoot.expand();
	} else {
		west.expandAll();
	}

	 $j("#center-panel > div > div").resize(resizeFrame);
	 resizeFrame();


	
	
		viewDoc('20191114001273', '6958001', '1', '670', '4055', 'dart3.xsd');
	
	




//팝업 순서
//1.정정 관련 보고서
//2.핵심투자 관련 보고서


	
	
	
	
	


alertInvestNotice('20191114001273', '6958001', '20191114001073', '0', '11013' ,'dart3.xsd');
alertSmallPublicOffering('N');






	window.focus();

	try {
		opener.loadHistoryContents();
	} catch (e) {
		;
	}
}

var currentDocValues = {rcpNo: "", dcmNo: "", eleId: "", offset: "", length: "", dtd: ""};

function replaceHtml(el, html) {
	var oldEl = (typeof el === "string" ? document.getElementById(el) : el);
	/*@cc_on // Pure innerHTML is slightly faster in IE
		oldEl.innerHTML = html;
		return oldEl;
	@*/
	var newEl = oldEl.cloneNode(false);
	newEl.innerHTML = html;
	oldEl.parentNode.replaceChild(newEl, oldEl);
	/* Since we just removed the old element from the DOM, return a reference
	to the new element, which can be used to restore variable references. */
	return newEl;
};

function viewDoc(rcpNo, dcmNo, eleId, offset, length, dtd) {
	currentDocValues.rcpNo = rcpNo;
	currentDocValues.dcmNo = dcmNo;
	currentDocValues.eleId = eleId;
	currentDocValues.offset = offset;
	currentDocValues.length = length;
	currentDocValues.dtd = dtd;
	var params = "";
	params += "?rcpNo=" + rcpNo;
	params += "&dcmNo=" + dcmNo;
	if (eleId != null)
		params += "&eleId=" + eleId;
	if (offset != null)
		params += "&offset=" + offset;
	if (length != null)
		params += "&length=" + length;
	params += "&dtd=" + dtd;
	document.getElementById("ifrm").src = "/report/viewer.do" + params;
}

function printDoc() {
	ifrm.focus();
	window.print();
}

function reloadDoc() {
	viewDoc(currentDocValues.rcpNo, currentDocValues.dcmNo, currentDocValues.eleId, currentDocValues.offset, currentDocValues.length, currentDocValues.dtd);
}

function changeDoc(val) {
//	window.location.replace("/dsaf001/main.do?" + val);
//	alert("/dsaf001/main.do?" + val);
	document.fmove.action="/dsaf001/main.do?" + val;
	document.fmove.submit();
}

function changeFamily(val) {
	if (val=="null") return;
	changeDoc(val);
	$j("#att").val("null");
	$j("#ref").val("null");
}

function changeRef(val) {
	if (val=="null") return;
	changeDoc(val);
	$j("#att").val("null");
	$j("#family").val("null");
}

function changeAtt(val) {
	if (val=="null") return;
	changeDoc(val);
	$j("#ref").val("null");
	$j("#family").val("null");
}

/**
 * PDF 다운로드 : 서비스 중지 메시지 보이기
 */
function openPdfDownload(rcpNo, dcmNo) {
	 
		
		var size = getOpenSize(530, 480);
		var url = "/pdf/download/main.do?rcp_no=" + rcpNo + "&dcm_no=" + dcmNo;
		var win = window.open( url, 'DOWNLOAD', ''+size+', resizable=no, status=no, scrollbars=yes');
		if(win == null){
			alert("팝업차단을 해제해주세요.");
		}else{
			win.focus();
		}
		
		
	
}

/**
 * 사용자가 '예'를 선택했을 때 최종문서로 이동한다.
 */
function confirmFinalDoc(bnt){

	var rcp_no = "" ;
	if(bnt=='yes'){
		
			
				if(rcp_no == ""){
					rcp_no = 20191114001273;
				}
			
		
		changeFamily("rcpNo="+rcp_no);
	}
}

function scrapFeed(cik) {
    var rssValue = "";
    if(cik == "today"){
    	rssValue = "http://dart.fss.or.kr/api/todayRSS.xml";
    }
    else{
    	rssValue = "http://dart.fss.or.kr/api/companyRSS.xml?crpCd=" + cik;
    }

    if(window.clipboardData){  // IE처리
    	if (window.clipboardData.setData("Text", rssValue)) {
        	alert('클립보드에 복사되었습니다');
    	}
    }
    else {                     // 비IE 처리
    	window.prompt ("Ctrl+C를 눌러 클립보드로 복사하세요", rssValue);
    }
}

	</script>
</head>
<body class="viewer">
<form action="" method="post" name="fmove">
<input name="goAction" style="display:none" type="submit"/>
</form>
<div class="header" id="north" style="width:1024px;"> <!-- -->
<div class="view_tit">
<p class="view_logo">
<img alt="DART" src="/images/common/viewer_logo.gif"/>
</p>
<p> <!--  class="viewTitle" -->
<!-- 회사명 표시  : 시장구분 이미지 : 회사명 : IR연결 아이콘-->
<img alt="코스피" src="/images/pop/ico_kospi.gif"/>
<a href="#companyview" onclick="openCorpInfo('00126380');return false;" style="cursor:pointer;" title="삼성전자">삼성전자</a><!-- 스타일 수정필요.  -->
</p>
</div>
<div class="view_search">
<!-- 본문선택 -->
<!-- <form method="" action=""> -->
<fieldset>
<legend>공시뷰어 검색</legend>
<p>
<label for="family">본문</label>
<select id="family" onchange="changeFamily(this.value)">
<option value="null">
								+본문선택+
							</option>
<option selected="" title="분기보고서" value="rcpNo=20191114001273">
									
										2019.11.14 
									
									
									분기보고서
								</option>
</select>
</p>
<p class="f_none">
<!-- 첨부문서 -->
<label for="att">첨부</label>
<select class="with" id="att" onchange="changeAtt(this.value)">
<option value="null">
								+첨부선택+
							</option>
<option value="rcpNo=20191114001273&amp;dcmNo=6958003">
									
									2019.11.14 
									
										
										
									
									분기검토보고서
								</option>
<option value="rcpNo=20191114001273&amp;dcmNo=6958004">
									
									2019.11.14 
									
										
										
									
									분기연결검토보고서
								</option>
</select>
</p>
</fieldset>
<!-- </form> -->
<ul><!-- 버튼 : 이미지보기, 다운로드, 인쇄, 검색결과로  -->
<li>
<a href="#download" onclick="openPdfDownload('20191114001273', '6958001'); return false;"><img alt="다운로드" src="/images/common/viewer_down.gif" style="cursor:pointer;" title="다운로드"/></a>
</li>
<li>
<a href="#priter" onclick="printDoc(); return false;"><img alt="인쇄" src="/images/common/viewer_print.gif" style="cursor:pointer;" title="인쇄"/></a>
</li>
<li>
<a href="#close" onclick="window.close();"><img alt="검색결과로" src="/images/common/viewer_result.gif" style="cursor:pointer;" title="검색결과로"/></a>
</li>
</ul>
</div>
</div><!-- end of header -->
<div id="center">
<div id="mobileScroll">
<iframe frameborder="0" id="ifrm" scrolling="auto" style="width: 100%;margin: 0; padding: 0" title="tailFrame"></iframe>
</div>
</div>
<div id="pdfStopMsg" style="position: absolute; left: -1000px; top: -1000px; visibility: hidden;"></div>
<script type="text/javascript">
function openReportViewerEng(rcpNo, dcmNo){
	var url = "/dsbh001/main.do?rcpNo=" + rcpNo;
	if (dcmNo) {
		//<![CDATA[
		url += "&dcmNo=" + dcmNo;
		//]]>
	}

	var size = getOpenSize(1024, 768);
	window.open(url, rcpNo, size+",resizable=yes");
	$j("#r_"+rcpNo).css("color","#cd0093");
}

var win_num = 1;
function openReportViewer(rcpNo, dcmNo){
	var url = "/dsaf001/main.do?rcpNo=" + rcpNo;
	if (dcmNo) {
		//<![CDATA[
		url += "&dcmNo=" + dcmNo;
		//]]>
	}

	var size = getOpenSize(1024, 768);
	window.open(url, rcpNo+win_num, size+",resizable=yes");
	win_num++;
	if (dcmNo) {
		$j("#r_"+rcpNo+dcmNo).css("color","#cd0093");
	}else{
		$j("#r_"+rcpNo).css("color","#cd0093");
	}
}

function openReportViewerDetail(rcpNo, dcmNo, detailYn, eleId, offSet, length, dtd){
	var url = "/dsaf001/main.do?rcpNo=" + rcpNo;
	if (dcmNo) {
		url += "&dcmNo=" + dcmNo;
	}
	if (detailYn) {
		url += "&detailYn=" + detailYn;
	}
	if (eleId) {
		url += "&eleId=" + eleId;
	}
	if (offSet) {
		url += "&offSet=" + offSet;
	}
	if (length) {
		url += "&length=" + length;
	}
	if (dtd) {
		url += "&dtd=" + dtd;
	}

	var size = getOpenSize(1024, 768);
	window.open(url, rcpNo+win_num, size+",resizable=yes");
	win_num++;
	if (dcmNo) {
		$j("#r_"+rcpNo+dcmNo).css("color","#cd0093");
	}else{
		$j("#r_"+rcpNo).css("color","#cd0093");
	}
}

function openReportViewerDetail(rcpNo, page, toc, gubun){
	//page = 4(재무), 5(사업보고서)
	//page = 4, toc = 0(재무-연결) | 5(재무-개별), gubun = J(주석)
	//page = 5, gubun = A부터

	var url = "/dsext004/viewer.do?rcpNo=" + rcpNo + "&page=" + page + "&selectToc=" + toc + "&gubun=" + gubun;

	var size = getOpenSize(1024, 768);
	window.open(url, rcpNo+win_num, size+",resizable=yes");
	win_num++;
	$j("#r_"+rcpNo).css("color","#cd0093");
}

/*
 * 통합검색용 문서뷰어 팝업 호출
 * nrseo 2010. 10. 20
 * 2016.09.30 jsp에 직접 명시 하여 미사용(DSAB007A.jsp)
 * - a:visited 미작동으로 인한 조치
 */
function openReportViewerKeyword(rcpNo, dcmNo, keyword){
	var url = "/dsaf001/main.do?rcpNo=" + rcpNo;
	if (dcmNo) {
		//<![CDATA[
		url += "&dcmNo=" + dcmNo;
		//]]>
	}
	//<![CDATA[
	url += "&keyword=" + encodeURIComponent(keyword);
	//]]>

	var size = getOpenSize(1024, 768);
	window.open(url, rcpNo+win_num, size+",resizable=yes");
	win_num++;
	if (dcmNo) {
		$j("#r_"+rcpNo+dcmNo).css("color","#cd0093");
	}else{
		$j("#r_"+rcpNo).css("color","#cd0093");
	}
}

/*
 * 메인화면 전용
 */
function openReportViewerMain(rcpNo, dcmNo){
	var url = "/dsaf001/main.do?rcpNo=" + rcpNo;
	if (dcmNo) {
		//<![CDATA[
		url += "&dcmNo=" + dcmNo;
		//]]>
	}

	var size = getOpenSize(1024, 768);
	window.open(url, rcpNo+win_num, size+",resizable=yes");
	win_num++;
}

function openXbrlViewer(rcpNo, dcmNo){
	var url = "/dsbh002/main.do?rcpNo=" + rcpNo;
	if (dcmNo) {
		//<![CDATA[
		url += "&dcmNo=" + dcmNo;
		//]]>
	}

	var size = getOpenSize(1024, 768);
	window.open(url, rcpNo+win_num, size+",resizable=yes");
	win_num++;
	/* j("#r_"+rcpNo)가
	   ie8,9,11에서 클릭시 css변경 미처리로 위치 변경
	   ie10 버전에서 미작동으로 본문에 직접 반영(DSBD001A.jsp, DSBD002A.jsp)	 */
	//$j("#r_"+rcpNo).css("color","#cd0093");
}

function openHelp(menu) {
	var url = "/guide/main.jsp";
	if (menu) {
		url += "?menu=" + menu;
	}
	var size = getOpenSize(800, 760);
	var win = window.open(url, "DartGuide", size + ",resizable=0,scrollbars=0");
	if (win == null) {
		alert("팝업차단을 해제하시기 바랍니다.");
	} else {
		win.focus();
	}
}

</script>
<form action="/dsaf001/main.do" method="post" name="reportForm">
<input name="rcpNo" type="hidden"/>
<input name="dcmNo" type="hidden"/>
<input name="keyword" type="hidden"/>
<input style="display:none;" type="submit"/>
</form>
</body>
</html>