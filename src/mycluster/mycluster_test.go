package mycluster

import (
	//"os/exec"
	//"syscall"
	"fmt"
	"strconv"
	"testing"
	//"os"
	"sync"
	"time"
)

func actualtest(servermainstruct []Servermainstruct) bool {

	retval := false
	i, j, k := 0, 0, 0

	for i = 0; i < len(servermainstruct); i++ {

		for j = 0; j < len(servermainstruct); j++ {

			for k = 0; k < 2; k++ {
				if servermainstruct[i].Actuallysend[j][k] != servermainstruct[j].Actuallyrecvd[i][k] {
					break
				}
			}
		}
	}
	if i == len(servermainstruct) && j == len(servermainstruct) && k == 2 {
		retval = true
	}

	return retval
}

/*func TestLargesizebroadcast(t *testing.T) {
	numberofservers:=5

	//numberofservers,_=strconv.Atoi(numberofservers)
	wg := new(sync.WaitGroup)
    	wg.Add(numberofservers)

	sizeofinchan:=20
	sizeofoutchan:=20

	//ffnm := "serverlist" + id + ".xml"
	//myservermainstruct := New(id,sizeofinchan,sizeooutchan,ffnm)
	//go myservermainstruct.Sendtooutbox(id, noofmessagestosend, numberofservers)
	myservermainstruct:=make([]Servermainstruct,numberofservers)

	//please start this broadcast string with string "BROADCAST" to get bash output right
	broadcastmsg:="BROADCAST mQINBE6txjkBEACv71ipnEvKO79OSiGOOa0eiphv1DolHZ3JsgJ0BOtsZqXjy9iwfRR81iyaDdRkE1PfGzi21PKoov7xhC32TeGSHtU2kittk3v9xPUz3aBYtXi1bbdnxqyvPJ7MYi4vfazLVrCVubJ5Fc/"

	//please start this broadcast string with string "hello there" to get bash output right
	peermsg:="hello there mQINBE6txjkBEACv71ipnEvKO79OSiGOOa0eiphv1DolHZ3JsgJ0BOtsZqXjy9iwfRR81iyaDdRkE1PfGzi21PKoov7xhC32TeGSHtU2kittk3v9xPUz3aBYtXi1bbdnxqyvPJ7MYi4vfazLVrCVubJ5Fc/vZhCOhhJA3YNA8qp97XD0JNICTxp5yRzRm2gCRY0nQFhYtKlUy8wYfvBjq3iMZiDm3/mel3fitBhOLlAdYsUt9XIQy"
	noofbroadcastmsgs:=10000
	noofpeermsgs:=10000

	noofmessagestosend:=noofbroadcastmsgs+noofpeermsgs


	delaybeforeconn:=3000 * time.Millisecond


	for i:=0;i<numberofservers;i++ {
		id:=strconv.Itoa(i)
		ffnm := "serverlist" + id + ".xml"
		myservermainstruct[i]= New(id,sizeofinchan,sizeofoutchan,ffnm,delaybeforeconn)
		go Serverfunc(myservermainstruct[i],id,noofmessagestosend,numberofservers,sizeofinchan,sizeofoutchan,wg)
		go myservermainstruct[i].Sendtooutbox(id, noofmessagestosend,broadcastmsg,peermsg,noofbroadcastmsgs,noofpeermsgs,numberofservers)
	}

	wg.Wait()

	if(actualtest(myservermainstruct)) {
		fmt.Println("My 1st Test Passed")
	} else {
		fmt.Println("My 1st Test is not Passed")
	}


}*/

/*func TestAllbroadcasts(t *testing.T) {
	numberofservers:=5
	//numberofservers,_=strconv.Atoi(numberofservers)
	wg := new(sync.WaitGroup)
    	wg.Add(numberofservers)

	sizeofinchan:=20
	sizeofoutchan:=20

	//ffnm := "serverlist" + id + ".xml"
	//myservermainstruct := New(id,sizeofinchan,sizeooutchan,ffnm)
	//go myservermainstruct.Sendtooutbox(id, noofmessagestosend, numberofservers)
	myservermainstruct:=make([]Servermainstruct,numberofservers)

	//please start this broadcast string with string "BROADCAST" to get bash output right
	broadcastmsg:="BROADCAST"

	//please start this broadcast string with string "hello there" to get bash output right
	peermsg:="hello there"
	noofbroadcastmsgs:=10000
	noofpeermsgs:=0

	noofmessagestosend:=noofbroadcastmsgs+noofpeermsgs


	delaybeforeconn:=3000 * time.Millisecond


	for i:=0;i<numberofservers;i++ {
		id:=strconv.Itoa(i)
		ffnm := "serverlist" + id + ".xml"
		myservermainstruct[i]= New(id,sizeofinchan,sizeofoutchan,ffnm,delaybeforeconn)
		go Serverfunc(myservermainstruct[i],id,noofmessagestosend,numberofservers,sizeofinchan,sizeofoutchan,wg)
		go myservermainstruct[i].Sendtooutbox(id, noofmessagestosend,broadcastmsg,peermsg,noofbroadcastmsgs,noofpeermsgs,numberofservers)
	}

	wg.Wait()


	if(actualtest(myservermainstruct)) {
		fmt.Println("My 2nd Test Passed")
	} else {
		fmt.Println("My 2nd Test is not Passed")
	}


}*/

/*func TestRoundrobbin(t *testing.T) {

	numberofservers:=5

	wg := new(sync.WaitGroup)
    	wg.Add(numberofservers)

	sizeofinchan:=20
	sizeofoutchan:=20

	//ffnm := "serverlist" + id + ".xml"
	//myservermainstruct := New(id,sizeofinchan,sizeooutchan,ffnm)
	//go myservermainstruct.Sendtooutbox(id, noofmessagestosend, numberofservers)
	myservermainstruct:=make([]Servermainstruct,numberofservers)

	//please start this broadcast string with string "BROADCAST" to get bash output right
	broadcastmsg:="BROADCAST"

	//please start this broadcast string with string "hello there" to get bash output right
	peermsg:="hello there"
	noofbroadcastmsgs:=10000
	noofpeermsgs:=10000

	noofmessagestosend:=noofbroadcastmsgs+noofpeermsgs


	delaybeforeconn:=3000 * time.Millisecond


	for i:=0;i<numberofservers;i++ {
		id:=strconv.Itoa(i)
		ffnm := "serverlist" + id + ".xml"
		myservermainstruct[i]= New(id,sizeofinchan,sizeofoutchan,ffnm,delaybeforeconn)
		go Serverfunc(myservermainstruct[i],id,noofmessagestosend,numberofservers,sizeofinchan,sizeofoutchan,wg)
		go myservermainstruct[i].Sendtooutbox3(id, noofmessagestosend,broadcastmsg,peermsg,noofbroadcastmsgs,noofpeermsgs,numberofservers)
	}

	wg.Wait()

	if(actualtest(myservermainstruct)) {
		fmt.Println("My 3rd Test Passed")
	} else {
		fmt.Println("My 3rd Test is not Passed")
	}





}*/

func TestBroadcastwithonenotstarted(t *testing.T) {

	numberofservers := 5
	//numberofservers,_=strconv.Atoi(numberofservers)
	wg := new(sync.WaitGroup)
	wg.Add(numberofservers)

	sizeofinchan := 20
	sizeofoutchan := 20

	//ffnm := "serverlist" + id + ".xml"
	//myservermainstruct := New(id,sizeofinchan,sizeooutchan,ffnm)
	//go myservermainstruct.Sendtooutbox(id, noofmessagestosend, numberofservers)
	myservermainstruct := make([]Servermainstruct, numberofservers)

	//please start this broadcast string with string "BROADCAST" to get bash output right
	broadcastmsg := "BROADCAST"

	//please start this broadcast string with string "hello there" to get bash output right
	peermsg := "hello there"
	noofbroadcastmsgs := 10000
	noofpeermsgs := 10000

	noofmessagestosend := noofbroadcastmsgs + noofpeermsgs

	delaybeforeconn := 3000 * time.Millisecond

	i := 0
	for i = 0; i < numberofservers; i++ {
		id := strconv.Itoa(i)
		ffnm := "serverlist" + id + ".xml"
		myservermainstruct[i] = New(id, sizeofinchan, sizeofoutchan, ffnm, delaybeforeconn)
		go Serverfunc(myservermainstruct[i], id, noofmessagestosend, numberofservers, sizeofinchan, sizeofoutchan, wg)
		go myservermainstruct[i].Sendtooutbox(id, noofmessagestosend, broadcastmsg, peermsg, noofbroadcastmsgs, noofpeermsgs, numberofservers)
		if i == numberofservers-2 {
			time.Sleep(10000 * time.Millisecond)
		}
	}

	wg.Wait()

	if actualtest(myservermainstruct) {
		fmt.Println("My 4th Test Passed")
	} else {
		fmt.Println("My 4th Test is not Passed")
	}
}
