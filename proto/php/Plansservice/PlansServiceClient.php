<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Plansservice;

/**
 */
class PlansServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Plansservice\PlanRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function CreatePlan(\Plansservice\PlanRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/plansservice.PlansService/CreatePlan',
        $argument,
        ['\Plansservice\PlanCreated', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \Plansservice\GetPlanRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function GetPlan(\Plansservice\GetPlanRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/plansservice.PlansService/GetPlan',
        $argument,
        ['\Plansservice\Plan', 'decode'],
        $metadata, $options);
    }

}
