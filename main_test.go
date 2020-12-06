package main

import (
   "testing"
   "os"
)

func TestHomePage(t *testing.T){
 var (
		links    []string 
		activelinks []string
	)
	
   links,activelinks = crawlUrl(os.Getenv("APP_BASE_URL"))
   
   if len(links) != len(activelinks){
	    t.Errorf("Hidden URL(s) got exposed in code which is subject to vulnerable")
   } 
   
}

func TestProductPage(t *testing.T){
 var (
		links    []string 
		activelinks []string
	)
	
   links,activelinks = crawlUrl(os.Getenv("APP_BASE_URL")+"/products")
   
   if len(links) != len(activelinks){
	    t.Errorf("Hidden URL(s) got exposed in code which is subject to vulnerable")
   } 
   
}


func TestCustomersPage(t *testing.T){
 var (
		links    []string 
		activelinks []string
	)
	
   links,activelinks = crawlUrl(os.Getenv("APP_BASE_URL")+"/customers")
   
   if len(links) != len(activelinks){
	    t.Errorf("Hidden URL(s) got exposed in code which is subject to vulnerable")
   } 
   
}

func TestRegisterPage(t *testing.T){
 var (
		links    []string 
		activelinks []string
	)
	
   links,activelinks = crawlUrl(os.Getenv("APP_BASE_URL")+"/admin/register")
   
   if len(links) != len(activelinks){
	    t.Errorf("Hidden URL(s) got exposed in code which is subject to vulnerable %s", links[len(activelinks):len(links)])
		
   } 
   
}