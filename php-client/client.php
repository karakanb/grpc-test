<?php

require dirname(__FILE__) . '/vendor/autoload.php';

use Plansservice\PlansServiceClient;
use Plansservice\GetPlanRequest;

$client = new PlansServiceClient('localhost:8000', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

$request = new GetPlanRequest();
$request->setID('plan_id');

while (true) {

    /** @var Plansservice\Plan $response  */
    list($response, $status) = $client->GetPlan($request)->wait();
    echo sprintf("Fetched a plan | ID: %s - sku: %s - created at: %s" . PHP_EOL, $response->getID(), $response->getSku(), $response->getCreatedAt());

    // print_r([
    //     'id' => $response->getID(),
    //     'customerUUID' => $response->getCustomerUUID(),
    //     'sku' => $response->getSku(),
    //     'created_at' => $response->getCreatedAt(),
    //     'updated_at' => $response->getUpdatedAt(),

    // ]);

    sleep(2);
}
