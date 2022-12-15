- Решение заключается в следующем: мы строим НКА, потом для каждого запроса идем по автомату с помощью DFS, параллельно помечая на ребрах длину префикса, который мы прочитали, проходя через эти ребра. Нам это нужно, чтобы избежать зацикливания из-за эпсилон-переходов*. Понятно, что если мы очередной раз проходим по ребру, по которому раньше уже проходили и "ничего не изменилось", то есть не изменилась длина префикса, который мы можем прочитать, то нам следует завершиться. 
- Мы найдем правильный ответ чисто по определениям НКА и DFS-а: возьмем максимальный префикс, который мы можем прочитать. Тогда существует путь, который соответствует этому префиксу. Возьмем кратчайший из всех путей. Напишем на каждом ребре длину префикса, который мы прочитали, проходя по этому ребру (если по какому-то ребру проходили несколько раз, то напишем несколько чисел). Понятно, что мы не могли написать два одинаковых числа на каком-то ребре, тк если есть такой "цикл", то мы можем его просто выкинуть, т.к. он ни на что по сути не влияет. Но получившийся путь обязательно будет пройден, как раз из-за корректности DFS-а. Таким образом, мы прочитаем максимальный префикс слова, который лежит в заданном языке.

*мы не будем использовать классический массив used, а будем проверять очередное просматриваемое ребро, а именно проверять, какой префикс мы прочитали, проходя по этому ребру в прошлый раз: если он был таким же по длине -- то не идем по этому ребру, если он не был таким же -- то идем. 