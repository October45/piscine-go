STREETS=$(grep "Annabel" ../the-final-cl-test/mystery/people | rev | cut -d$'\t' -f1 | rev | cut -d',' -f1 | sed 's/ /_/g')
LINES=$(grep "Annabel" ../the-final-cl-test/mystery/people | rev | cut -d$' ' -f1 | rev)

STREETS=($STREETS)
LINES=($LINES)

for index in ${!STREETS[@]}
do
    interview_nr=$(cat ../the-final-cl-test/mystery/streets/${STREETS[$index]} | awk -v line=${LINES[$index]} 'NR == line {print $3}' | tr -d "#")
    eval 'interview_nr'"$index"'=interview-"$interview_nr"'
done
echo $interview_nr0

cat "../the-final-cl-test/mystery/interviews/${interview_nr0}"