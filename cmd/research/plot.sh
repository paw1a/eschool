#!/opt/homebrew/bin/gnuplot -persist

reset

set terminal svg size 1280, 720 font "Helvetica,18"
set output 'plot3.svg'
set size ratio 0.56
set pointsize 0.9

set key reverse Left
set xlabel "Количество записей в таблице уроков, шт."
set ylabel "Время, мс"
set grid

plot "./data.txt" using 1:2 with linespoints title 'Проверка курса на стороне приложения' pt 9,\
"./data.txt" using 1:3 with linespoints title 'Проверка курса на стороне базы данных' pt 10,\
