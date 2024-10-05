/*  This query demonstrates the typical level of complexity that I may produce in a private, enterprise environment.
    I want sql to show up on my GitHub profile as a language skill since i do not typically use it personally & therefor would not show up. */

--EXPLAIN ANALYZE

WITH contract_info AS (
    SELECT
        c.contract_id,
        0 AS "person_id",
        c.contract_holder_first_name AS "first_name",
        c.contract_holder_last_name AS "last_name",
        c.department
    FROM
        contractor c
),

dependent_info AS (
    SELECT
        d.contract_id,
        d.person_id,
        d.dependent_first_name AS "first_name",
        d.dependent_last_name AS "last_name",
        d.department
    FROM 
        dependent d
)

SELECT
    ci.contract_id,
    LPAD(COALESCE(di."person_id", 0)::text, 2, '0') AS "person_id",
    COALESCE(ci."first_name", di."first_name") AS "first_name",
    COALESCE(ci."last_name", di."last_name") AS "last_name",
    CASE 
        WHEN ci.department = 1 THEN 'DISC'
        WHEN ci.department = 2 THEN 'FREE'
        WHEN ci.department = 3 THEN 'FULL'
        ELSE 'UNKNOWN'
    END AS "department",
    c.coverage_status,
    c.effective_date::date AS "effective_date",
    c.expiration_date::date AS "expiration_date"
FROM
    contract_info ci
LEFT JOIN
    dependent_info di ON ci.contract_id = di.contract_id
LEFT JOIN
    coverage c ON ci.contract_id = c.contract_id AND COALESCE(di.person_id, 0) = c.person_id
ORDER BY
    ci.contract_id, "person_id";