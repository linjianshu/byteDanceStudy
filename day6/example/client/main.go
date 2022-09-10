package main

import (
		"context"
			"example/kitex_gen/api"
				"example/kitex_gen/api/echo"
					"github.com/cloudwego/kitex/client"
						"github.com/cloudwego/kitex/client/callopt"
							"log"
								"time"
							)

							func main() {
									c, err := echo.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
										if err != nil {
													log.Fatal(err)
														}
															req := &api.Request{Message: "my request"}
																resp, err := c.Echo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
																	if err != nil {
																				log.Fatal(err)
																					}
																						log.Println(resp)
						addReq:=&api.AddRequest{First:512,Second:512}
						addResp,err:=c.Add(context.Background(),addReq)
						if err!=nil{
							log.Fatal(err)
						}
						log.Println(addResp)
						}

