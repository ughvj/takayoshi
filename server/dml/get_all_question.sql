select q.id, q.statement, q.category, qo.correct_choice, qo.correct_order, qo.name_kanji as genkun_name, qo.src as genkun_src
from questions as q
inner join (
    select qo.question_id, qo.correct_choice, qo.correct_order, g.name_kanji, g.src
    from question_options as qo
    inner join genkuns as g
    on qo.genkun_id = g.id
) as qo
on q.id = qo.question_id
order by q.id asc;
