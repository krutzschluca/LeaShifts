<?php

// Dummy shift data (later will be fetched from database)
$shiftPlanData = [
    "Monday" => [
        "08:00" => "John",
        "09:00" => "John",
        "10:00" => "John",
        "11:00" => "John",
        "12:00" => "John",
        "13:00" => "John",
        "14:00" => "John",
        "15:00" => "John",
        "16:00" => "John",
        "17:00" => "John"
    ],
    "Tuesday" => [
        "08:00" => "David",
        "09:00" => "David",
        "10:00" => "David",
        "11:00" => "David",
        "12:00" => "Emily",
        "13:00" => "Emily",
        "14:00" => "Emily",
        "15:00" => "Emily",
        "16:00" => "Emily",
        "17:00" => "Emily"
    ],
    "Wednesday" => [
        "10:00" => "Bob",
        "11:00" => "Bob",
        "12:00" => "Bob",
        "13:00" => "Bob",
        "14:00" => "Bob",
        "15:00" => "Bob",
        "16:00" => "Bob",
        "17:00" => "Bob"
    ],
    "Thursday" => [

        "12:00" => "David",
        "13:00" => "David",
        "14:00" => "David",
        "15:00" => "David",
        "16:00" => "David",
        "17:00" => "David"
    ],
    "Friday" => [
        "08:00" => "Alice",
        "09:00" => "Alice",
        "10:00" => "Alice",
        "11:00" => "Alice",
    ],
    "Saturday" => [
        "08:00" => "Emily",
        "09:00" => "Emily",
        "10:00" => "Emily",
        "11:00" => "Emily",
        "12:00" => "Emily",
        "13:00" => "Emily",
        "14:00" => "Emily",
        "15:00" => "Emily",
        "16:00" => "Emily",
        "17:00" => "Emily"
    ],
    "Sunday" => [
        "08:00" => "David",
        "09:00" => "John",
        "10:00" => "Alice",
        "11:00" => "Bob",
        "12:00" => "Emily",
        "13:00" => "David",
        "14:00" => "John",
        "15:00" => "Alice",
        "16:00" => "Bob",
        "17:00" => "Emily"
    ],
];


header('Content-Type: application/json');

// Output the JSON data
echo json_encode($shiftPlanData);
