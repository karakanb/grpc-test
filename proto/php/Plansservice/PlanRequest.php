<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: plans_service.proto

namespace Plansservice;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>plansservice.PlanRequest</code>
 */
class PlanRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string customerUUID = 1;</code>
     */
    protected $customerUUID = '';
    /**
     * Generated from protobuf field <code>string sku = 2;</code>
     */
    protected $sku = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $customerUUID
     *     @type string $sku
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\PlansService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string customerUUID = 1;</code>
     * @return string
     */
    public function getCustomerUUID()
    {
        return $this->customerUUID;
    }

    /**
     * Generated from protobuf field <code>string customerUUID = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setCustomerUUID($var)
    {
        GPBUtil::checkString($var, True);
        $this->customerUUID = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string sku = 2;</code>
     * @return string
     */
    public function getSku()
    {
        return $this->sku;
    }

    /**
     * Generated from protobuf field <code>string sku = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setSku($var)
    {
        GPBUtil::checkString($var, True);
        $this->sku = $var;

        return $this;
    }

}
