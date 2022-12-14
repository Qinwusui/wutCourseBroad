package main

type CourseData struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}
type KckbData struct {
	ID      string `json:"id"`
	Xnxq    string `json:"xnxq"`
	Tid     string `json:"tid"`
	Type    int    `json:"type"`
	Xs      int    `json:"xs"`
	Rqxl    string `json:"rqxl"`
	Sfwc    int    `json:"sfwc"`
	Jxbid   string `json:"jxbid"`
	Jxbmc   string `json:"jxbmc"`
	Zctype  string `json:"zctype"`
	Kcmc    string `json:"kcmc"`
	Kcbh    string `json:"kcbh"`
	Zc      string `json:"zc"`
	Zcstr   string `json:"zcstr"`
	Croommc string `json:"croommc"`
	Tmc     string `json:"tmc"`
	Xqid    string `json:"xqid"`
	Xqmc    string `json:"xqmc"`
	Xingqi  int    `json:"xingqi"`
	Djc     int    `json:"djc"`
	Flag    int    `json:"flag"`
	Source  string `json:"source"`
	Pkid    string `json:"pkid"`
	Xq      string `json:"xq"`
	Bjdm    string `json:"bjdm"`
	Kcxz    string `json:"kcxz"`
	Xdxz    string `json:"xdxz"`
	Ksxs    string `json:"ksxs"`
}
type Data struct {
	KckbData []KckbData `json:"kckbData"`
}
type UserData struct {
	UserName string `json:"userName"`
	Pwd      string `json:"Pwd"`
}
