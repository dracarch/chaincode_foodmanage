package mockactions

import (
	"github.com/dormao/chaincode_foodmanage/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func RegTransporter(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.UnAuthRegisterTransporter),
		[]byte(TestPassword),
	})
	transporter_id = string(resp.GetPayload())
	return resp
}

func TransporterLogin(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.UnAuthLogin),
		[]byte(transporter_id),
		[]byte(TestPassword),
	})
	transporter_token = string(resp.GetPayload())
	return resp
}

func TransporterUpdateTemperature(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_UPDATE_TRANSPORT),
		[]byte(createCredentialsWithToken(transporter_token)),
		[]byte(transport_id),
		jsonEncode(&models.TransportDetails{
			Temperature: MockTransportTemperature,
		}),
	})
	return resp
}

func TransporterCancelTransport(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_CANCELTRANSPORT),
		[]byte(createCredentialsWithToken(transporter_token)),
		[]byte(transport_id),
	})
	return resp
}

func TransporterCompleteTransport(stub *shim.MockStub) peer.Response{
	resp := stub.MockInvoke(models.AllocateIdS(), [][]byte{
		[]byte(models.OPERATE_COMPLETE_TRANSPORT),
		[]byte(createCredentialsWithToken(transporter_token)),
		[]byte(transport_id),
	})
	return resp
}
