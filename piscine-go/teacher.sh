#! /bin/bash
interviewn=`grep -H "licen" interviews/* |grep "\"" | cut -f1 -d ":" | rev | cut -f1 -d "-" | rev`
interview="cat interviews/interview-$interviewn"
export interviewnum=$interviewn
echo $interviewnum
$interview
echo $MAIN_SUSPECT