@echo off

echo SHB

REM mkdir results\shb

REM echo AVRORA 1/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\avrora.log > .\results\shb\avrora.log
REM echo BATIK 2/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\batikNT.log > .\results\shb\batik.log
REM echo CRYPT 3/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\cryptNT.log > .\results\shb\crypt.log
REM echo H2 4/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\h2NT.log -n 100000000 > .\results\shb\h2.log
REM echo LUFACT 5/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\lufactNT.log > .\results\shb\lufact.log
REM echo LUSEARCH 6/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\lusearchNT.log  > .\results\shb\lusearch.log
REM echo MOLDYN 7/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\moldyn.log  > .\results\shb\moldyn.log
REM echo MONTE 8/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\montecarloNT.log  > .\results\shb\monte.log
REM echo PMD 9/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\pmdNT.log  > .\results\shb\pmd.log
REM echo SERIESC 10/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\seriesCNT.log  > .\results\shb\seriesc.log
REM echo SOR 11/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\sorNT.log  > .\results\shb\sor.log
REM echo SUNFLOW 12/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\sunflowNTABORTED.log -n 100000000 > .\results\shb\sunflow.log
REM echo XALAN 13/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\xalanFull.log > .\results\shb\xalan.log
REM echo TOMCAT 14/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\tomcatFull.log > .\results\shb\tomcat.log
REM echo Raytrace 15/15
REM go run main.go -mode shbDefault -parser javainc -trace E:\traces\raytracer.log  > .\results\shb\raytracer.log


REM echo W3POEEHBC

REM mkdir results\tsanEEHBC

REM echo AVRORA 1/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\avrora.log > .\results\tsanEEHBC\avrora.log
REM echo BATIK 2/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\batikNT.log > .\results\tsanEEHBC\batik.log
REM echo CRYPT 3/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\cryptNT.log  > .\results\tsanEEHBC\crypt.log
REM echo H2 4/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\h2NT.log -n 100000000 > .\results\tsanEEHBC\h2.log
REM echo LUFACT 5/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\lufactNT.log  > .\results\tsanEEHBC\lufact.log
REM echo LUSEARCH 6/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\lusearchNT.log  > .\results\tsanEEHBC\lusearch.log
REM echo MOLDYN 7/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\moldyn.log  > .\results\tsanEEHBC\moldyn.log
REM echo MONTE 8/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\montecarloNT.log  > .\results\tsanEEHBC\monte.log
REM echo PMD 9/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\pmdNT.log  > .\results\tsanEEHBC\pmd.log
REM echo SERIESC 10/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\seriesCNT.log  > .\results\tsanEEHBC\seriesc.log
REM echo SOR 11/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\sorNT.log  > .\results\tsanEEHBC\sor.log
REM echo SUNFLOW 12/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\sunflowNTABORTED.log -n 100000000 > .\results\tsanEEHBC\sunflow.log
REM echo XALAN 13/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\xalanFull.log  > .\results\tsanEEHBC\xalan.log
REM echo TOMCAT 14/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\tomcatFull.log > .\results\tsanEEHBC\tomcat.log
REM echo Raytrace 15/15
REM go run main.go -mode tsanEEHBC -parser javainc -trace E:\traces\raytracer.log  > .\results\tsanEEHBC\raytracer.log

echo tsanEEW1

mkdir results\tsanEEW1

echo AVRORA 1/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\avrora.log > .\results\tsanEEW1\avrora.log
echo BATIK 2/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\batikNT.log > .\results\tsanEEW1\batik.log
echo CRYPT 3/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\cryptNT.log  > .\results\tsanEEW1\crypt.log
echo H2 4/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\h2NT.log -n 100000000 > .\results\tsanEEW1\h2.log
echo LUFACT 5/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\lufactNT.log  > .\results\tsanEEW1\lufact.log
echo LUSEARCH 6/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\lusearchNT.log  > .\results\tsanEEW1\lusearch.log
echo MOLDYN 7/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\moldyn.log  > .\results\tsanEEW1\moldyn.log
echo MONTE 8/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\montecarloNT.log  > .\results\tsanEEW1\monte.log
echo PMD 9/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\pmdNT.log  > .\results\tsanEEW1\pmd.log
echo SERIESC 10/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\seriesCNT.log  > .\results\tsanEEW1\seriesc.log
echo SOR 11/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\sorNT.log  > .\results\tsanEEW1\sor.log
echo SUNFLOW 12/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\sunflowNTABORTED.log -n 100000000 > .\results\tsanEEW1\sunflow.log
echo XALAN 13/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\xalanFull.log  > .\results\tsanEEW1\xalan.log
echo TOMCAT 14/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\tomcatFull.log > .\results\tsanEEW1\tomcat.log
echo Raytrace 15/15
go run main.go -mode tsanEEW1 -parser javainc2 -trace E:\traces\raytracer.log  > .\results\tsanEEW1\raytracer.log


REM echo TSANEE

REM mkdir results\tsanee

REM echo AVRORA 1/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\avrora.log > .\results\tsanee\avrora.log
REM echo BATIK 2/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\batikNT.log > .\results\tsanee\batik.log
REM echo CRYPT 3/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\cryptNT.log  > .\results\tsanee\crypt.log
REM echo H2 4/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\h2NT.log -n 100000000 > .\results\tsanee\h2.log
REM echo LUFACT 5/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\lufactNT.log  > .\results\tsanee\lufact.log
REM echo LUSEARCH 6/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\lusearchNT.log  > .\results\tsanee\lusearch.log
REM echo MOLDYN 7/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\moldyn.log  > .\results\tsanee\moldyn.log
REM echo MONTE 8/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\montecarloNT.log  > .\results\tsanee\monte.log
REM echo PMD 9/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\pmdNT.log  > .\results\tsanee\pmd.log
REM echo SERIESC 10/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\seriesCNT.log  > .\results\tsanee\seriesc.log
REM echo SOR 11/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\sorNT.log  > .\results\tsanee\sor.log
REM echo SUNFLOW 12/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\sunflowNTABORTED.log -n 100000000 > .\results\tsanee\sunflow.log
REM echo XALAN 13/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\xalanFull.log  > .\results\tsanee\xalan.log
REM echo TOMCAT 14/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\tomcatFull.log > .\results\tsanee\tomcat.log
REM echo Raytrace 15/15
REM go run main.go -mode tsanEE -parser javainc -trace E:\traces\raytracer.log  > .\results\tsanee\raytracer.log

REM echo TSANEEWRD

REM mkdir results\tsaneewrd

REM echo AVRORA 1/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\avrora.log > .\results\tsaneewrd\avrora.log
REM echo BATIK 2/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\batikNT.log > .\results\tsaneewrd\batik.log
REM echo CRYPT 3/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\cryptNT.log  > .\results\tsaneewrd\crypt.log
REM echo H2 4/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\h2NT.log -n 100000000 > .\results\tsaneewrd\h2.log
REM echo LUFACT 5/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\lufactNT.log  > .\results\tsaneewrd\lufact.log
REM echo LUSEARCH 6/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\lusearchNT.log  > .\results\tsaneewrd\lusearch.log
REM echo MOLDYN 7/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\moldyn.log  > .\results\tsaneewrd\moldyn.log
REM echo MONTE 8/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\montecarloNT.log  > .\results\tsaneewrd\monte.log
REM echo PMD 9/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\pmdNT.log  > .\results\tsaneewrd\pmd.log
REM echo SERIESC 10/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\seriesCNT.log  > .\results\tsaneewrd\seriesc.log
REM echo SOR 11/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\sorNT.log  > .\results\tsaneewrd\sor.log
REM echo SUNFLOW 12/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\sunflowNTABORTED.log -n 100000000 > .\results\tsaneewrd\sunflow.log
REM echo XALAN 13/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\xalanFull.log  > .\results\tsaneewrd\xalan.log
REM echo TOMCAT 14/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\tomcatFull.log > .\results\tsaneewrd\tomcat.log
REM echo Raytrace 15/15
REM go run main.go -mode tsanEEWRD -parser javainc -trace E:\traces\raytracer.log  > .\results\tsaneewrd\raytracer.log


REM echo W3POEE

REM mkdir results\w3poee

REM echo AVRORA 1/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\avrora.log > .\results\w3poee\avrora.log
REM echo BATIK 2/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\batikNT.log > .\results\w3poee\batik.log
REM echo CRYPT 3/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\cryptNT.log  > .\results\w3poee\crypt.log
REM echo H2 4/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\h2NT.log -n 100000000 > .\results\w3poee\h2.log
REM echo LUFACT 5/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\lufactNT.log  > .\results\w3poee\lufact.log
REM echo LUSEARCH 6/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\lusearchNT.log  > .\results\w3poee\lusearch.log
REM echo MOLDYN 7/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\moldyn.log  > .\results\w3poee\moldyn.log
REM echo MONTE 8/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\montecarloNT.log  > .\results\w3poee\monte.log
REM echo PMD 9/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\pmdNT.log  > .\results\w3poee\pmd.log
REM echo SERIESC 10/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\seriesCNT.log  > .\results\w3poee\seriesc.log
REM echo SOR 11/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\sorNT.log  > .\results\w3poee\sor.log
REM echo SUNFLOW 12/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\sunflowNTABORTED.log -n 100000000 > .\results\w3poee\sunflow.log
REM echo XALAN 13/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\xalanFull.log  > .\results\w3poee\xalan.log
REM echo TOMCAT 14/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\tomcatFull.log > .\results\w3poee\tomcat.log
REM echo Raytrace 15/15
REM go run main.go -mode tsanEEWCP -parser javainc -trace E:\traces\raytracer.log  > .\results\w3poee\raytracer.log


REM echo W3PO

REM mkdir results\w3po

REM echo AVRORA 1/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\avrora.log > .\results\w3po\avrora.log
REM echo BATIK 2/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\batikNT.log > .\results\w3po\batik.log
REM echo CRYPT 3/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\cryptNT.log > .\results\w3po\crypt.log
REM echo H2 4/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\h2NT.log -n 100000000 > .\results\w3po\h2.log
REM echo LUFACT 5/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\lufactNT.log > .\results\w3po\lufact.log
REM echo LUSEARCH 6/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\lusearchNT.log  > .\results\w3po\lusearch.log
REM echo MOLDYN 7/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\moldyn.log  > .\results\w3po\moldyn.log
REM echo MONTE 8/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\montecarloNT.log  > .\results\w3po\monte.log
REM echo PMD 9/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\pmdNT.log  > .\results\w3po\pmd.log
REM echo SERIESC 10/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\seriesCNT.log  > .\results\w3po\seriesc.log
REM echo SOR 11/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\sorNT.log  > .\results\w3po\sor.log
REM echo SUNFLOW 12/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\sunflowNTABORTED.log -n 100000000 > .\results\w3po\sunflow.log
REM echo XALAN 13/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\xalanFull.log > .\results\w3po\xalan.log
REM echo TOMCAT 14/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\tomcatFull.log> .\results\w3po\tomcat.log
REM echo Raytrace 15/15
REM go run main.go -mode tsanEWCP -parser javainc -trace E:\traces\raytracer.log  > .\results\w3po\raytracer.log


REM echo FSSHBEE

REM mkdir results\fsshbee

REM echo AVRORA 1/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\avrora.log > .\results\fsshbee\avrora.log
REM echo BATIK 2/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\batikNT.log > .\results\fsshbee\batik.log
REM echo CRYPT 3/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\cryptNT.log > .\results\fsshbee\crypt.log
REM echo H2 4/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\h2NT.log -n 100000000 > .\results\fsshbee\h2.log
REM echo LUFACT 5/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\lufactNT.log > .\results\fsshbee\lufact.log
REM echo LUSEARCH 6/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\lusearchNT.log  > .\results\fsshbee\lusearch.log
REM echo MOLDYN 7/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\moldyn.log  > .\results\fsshbee\moldyn.log
REM echo MONTE 8/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\montecarloNT.log  > .\results\fsshbee\monte.log
REM echo PMD 9/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\pmdNT.log  > .\results\fsshbee\pmd.log
REM echo SERIESC 10/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\seriesCNT.log  > .\results\fsshbee\seriesc.log
REM echo SOR 11/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\sorNT.log  > .\results\fsshbee\sor.log
REM echo SUNFLOW 12/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\sunflowNTABORTED.log -n 100000000 > .\results\fsshbee\sunflow.log
REM echo XALAN 13/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\xalanFull.log > .\results\fsshbee\xalan.log
REM echo TOMCAT 14/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\tomcatFull.log> .\results\fsshbee\tomcat.log
REM echo Raytrace 15/15
REM go run main.go -mode fsshbee -parser javainc -trace E:\traces\raytracer.log  > .\results\fsshbee\raytracer.log

REM echo WCP

REM mkdir results\wcp

REM echo AVRORA 1/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\avrora.log > .\results\wcp\avrora.log
REM echo BATIK 2/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\batikNT.log > .\results\wcp\batik.log
REM echo CRYPT 3/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\cryptNT.log > .\results\wcp\crypt.log
REM echo H2 4/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\h2NT.log -n 100000000 > .\results\wcp\h2.log
REM echo LUFACT 5/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\lufactNT.log > .\results\wcp\lufact.log
REM echo LUSEARCH 6/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\lusearchNT.log  > .\results\wcp\lusearch.log
REM echo MOLDYN 7/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\moldyn.log  > .\results\wcp\moldyn.log
REM echo MONTE 8/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\montecarloNT.log  > .\results\wcp\monte.log
REM echo PMD 9/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\pmdNT.log  > .\results\wcp\pmd.log
REM echo SERIESC 10/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\seriesCNT.log  > .\results\wcp\seriesc.log
REM echo SOR 11/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\sorNT.log  > .\results\wcp\sor.log
REM echo SUNFLOW 12/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\sunflowNTABORTED.log -n 100000000 > .\results\wcp\sunflow.log
REM echo XALAN 13/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\xalanFull.log > .\results\wcp\xalan.log
REM echo TOMCAT 14/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\tomcatFull.log> .\results\wcp\tomcat.log
REM echo Raytrace 15/15
REM go run main.go -mode wcpDefault -parser javainc -trace E:\traces\raytracer.log  > .\results\wcp\raytracer.log

REM echo TSAN

REM mkdir results\tsan

REM echo AVRORA 1/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\avrora.log > .\results\tsan\avrora.log
REM echo BATIK 2/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\batikNT.log > .\results\tsan\batik.log
REM echo CRYPT 3/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\cryptNT.log > .\results\tsan\crypt.log
REM echo H2 4/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\h2NT.log -n 100000000 > .\results\tsan\h2.log
REM echo LUFACT 5/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\lufactNT.log > .\results\tsan\lufact.log
REM echo LUSEARCH 6/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\lusearchNT.log  > .\results\tsan\lusearch.log
REM echo MOLDYN 7/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\moldyn.log  > .\results\tsan\moldyn.log
REM echo MONTE 8/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\montecarloNT.log  > .\results\tsan\monte.log
REM echo PMD 9/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\pmdNT.log  > .\results\tsan\pmd.log
REM echo SERIESC 10/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\seriesCNT.log  > .\results\tsan\seriesc.log
REM echo SOR 11/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\sorNT.log  > .\results\tsan\sor.log
REM echo SUNFLOW 12/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\sunflowNTABORTED.log -n 100000000 > .\results\tsan\sunflow.log
REM echo XALAN 13/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\xalanFull.log > .\results\tsan\xalan.log
REM echo TOMCAT 14/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\tomcatFull.log > .\results\tsan\tomcat.log
REM echo Raytrace 15/15
REM go run main.go -mode tsan -parser javainc -trace E:\traces\raytracer.log  > .\results\tsan\raytracer.log

REM echo HB

REM mkdir results\ft

REM echo AVRORA 1/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\avrora.log > .\results\ft\avrora.log
REM echo BATIK 2/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\batikNT.log > .\results\ft\batik.log
REM echo CRYPT 3/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\cryptNT.log > .\results\ft\crypt.log
REM echo H2 4/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\h2NT.log -n 100000000 > .\results\ft\h2.log
REM echo LUFACT 5/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\lufactNT.log > .\results\ft\lufact.log
REM echo LUSEARCH 6/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\lusearchNT.log  > .\results\ft\lusearch.log
REM echo MOLDYN 7/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\moldyn.log  > .\results\ft\moldyn.log
REM echo MONTE 8/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\montecarloNT.log  > .\results\ft\monte.log
REM echo PMD 9/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\pmdNT.log  > .\results\ft\pmd.log
REM echo SERIESC 10/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\seriesCNT.log  > .\results\ft\seriesc.log
REM echo SOR 11/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\sorNT.log  > .\results\ft\sor.log
REM echo SUNFLOW 12/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\sunflowNTABORTED.log -n 100000000 > .\results\ft\sunflow.log
REM echo XALAN 13/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\xalanFull.log > .\results\ft\xalan.log
REM echo TOMCAT 14/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\tomcatFull.log > .\results\ft\tomcat.log
REM echo Raytrace 15/15
REM go run main.go -mode hbDefault -parser javainc -trace E:\traces\raytracer.log  > .\results\ft\raytracer.log




