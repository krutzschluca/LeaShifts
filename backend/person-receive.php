<?php

// Dummy people data (later will be fetched from database)
$people = ["John", "Alice", "Bob", "Emily", "David"];

header('Content-Type: application/json');

// Output the JSON data
echo json_encode(['people' => $people]);
?>
