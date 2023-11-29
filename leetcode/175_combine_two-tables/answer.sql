SELECT
    p.firstName,
    p.lastName,
    COALESCE(a.city, NULL) as city,
    a.state
  FROM Person p
    LEFT JOIN Address a ON p.personId = a.personId
