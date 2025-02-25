// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Splight.Splight
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import splight:index/server:Server [options] splight_server.&lt;name&gt; &lt;server_id&gt;
    /// ```
    /// </summary>
    [SplightResourceType("splight:index/server:Server")]
    public partial class Server : global::Pulumi.CustomResource
    {
        /// <summary>
        /// static config parameters of the routine
        /// </summary>
        [Output("configs")]
        public Output<ImmutableArray<Outputs.ServerConfig>> Configs { get; private set; } = null!;

        /// <summary>
        /// optional description to add details of the resource
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// environment variables for the server
        /// </summary>
        [Output("envVars")]
        public Output<ImmutableArray<Outputs.ServerEnvVar>> EnvVars { get; private set; } = null!;

        /// <summary>
        /// log level of the server
        /// </summary>
        [Output("logLevel")]
        public Output<string?> LogLevel { get; private set; } = null!;

        /// <summary>
        /// instance size
        /// </summary>
        [Output("machineInstanceSize")]
        public Output<string?> MachineInstanceSize { get; private set; } = null!;

        /// <summary>
        /// the name of the server to be created
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// id of the compute node where the server runs
        /// </summary>
        [Output("node")]
        public Output<string?> Node { get; private set; } = null!;

        /// <summary>
        /// ports of the server
        /// </summary>
        [Output("ports")]
        public Output<ImmutableArray<Outputs.ServerPort>> Ports { get; private set; } = null!;

        /// <summary>
        /// restart policy of the server
        /// </summary>
        [Output("restartPolicy")]
        public Output<string?> RestartPolicy { get; private set; } = null!;

        /// <summary>
        /// tags of the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ServerTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// [NAME-VERSION] the version of the hub server
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Server resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Server(string name, ServerArgs args, CustomResourceOptions? options = null)
            : base("splight:index/server:Server", name, args ?? new ServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Server(string name, Input<string> id, ServerState? state = null, CustomResourceOptions? options = null)
            : base("splight:index/server:Server", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/splightplatform",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Server resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Server Get(string name, Input<string> id, ServerState? state = null, CustomResourceOptions? options = null)
        {
            return new Server(name, id, state, options);
        }
    }

    public sealed class ServerArgs : global::Pulumi.ResourceArgs
    {
        [Input("configs")]
        private InputList<Inputs.ServerConfigArgs>? _configs;

        /// <summary>
        /// static config parameters of the routine
        /// </summary>
        public InputList<Inputs.ServerConfigArgs> Configs
        {
            get => _configs ?? (_configs = new InputList<Inputs.ServerConfigArgs>());
            set => _configs = value;
        }

        /// <summary>
        /// optional description to add details of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("envVars")]
        private InputList<Inputs.ServerEnvVarArgs>? _envVars;

        /// <summary>
        /// environment variables for the server
        /// </summary>
        public InputList<Inputs.ServerEnvVarArgs> EnvVars
        {
            get => _envVars ?? (_envVars = new InputList<Inputs.ServerEnvVarArgs>());
            set => _envVars = value;
        }

        /// <summary>
        /// log level of the server
        /// </summary>
        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        /// <summary>
        /// instance size
        /// </summary>
        [Input("machineInstanceSize")]
        public Input<string>? MachineInstanceSize { get; set; }

        /// <summary>
        /// the name of the server to be created
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// id of the compute node where the server runs
        /// </summary>
        [Input("node")]
        public Input<string>? Node { get; set; }

        [Input("ports")]
        private InputList<Inputs.ServerPortArgs>? _ports;

        /// <summary>
        /// ports of the server
        /// </summary>
        public InputList<Inputs.ServerPortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.ServerPortArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// restart policy of the server
        /// </summary>
        [Input("restartPolicy")]
        public Input<string>? RestartPolicy { get; set; }

        [Input("tags")]
        private InputList<Inputs.ServerTagArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.ServerTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ServerTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// [NAME-VERSION] the version of the hub server
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public ServerArgs()
        {
        }
        public static new ServerArgs Empty => new ServerArgs();
    }

    public sealed class ServerState : global::Pulumi.ResourceArgs
    {
        [Input("configs")]
        private InputList<Inputs.ServerConfigGetArgs>? _configs;

        /// <summary>
        /// static config parameters of the routine
        /// </summary>
        public InputList<Inputs.ServerConfigGetArgs> Configs
        {
            get => _configs ?? (_configs = new InputList<Inputs.ServerConfigGetArgs>());
            set => _configs = value;
        }

        /// <summary>
        /// optional description to add details of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("envVars")]
        private InputList<Inputs.ServerEnvVarGetArgs>? _envVars;

        /// <summary>
        /// environment variables for the server
        /// </summary>
        public InputList<Inputs.ServerEnvVarGetArgs> EnvVars
        {
            get => _envVars ?? (_envVars = new InputList<Inputs.ServerEnvVarGetArgs>());
            set => _envVars = value;
        }

        /// <summary>
        /// log level of the server
        /// </summary>
        [Input("logLevel")]
        public Input<string>? LogLevel { get; set; }

        /// <summary>
        /// instance size
        /// </summary>
        [Input("machineInstanceSize")]
        public Input<string>? MachineInstanceSize { get; set; }

        /// <summary>
        /// the name of the server to be created
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// id of the compute node where the server runs
        /// </summary>
        [Input("node")]
        public Input<string>? Node { get; set; }

        [Input("ports")]
        private InputList<Inputs.ServerPortGetArgs>? _ports;

        /// <summary>
        /// ports of the server
        /// </summary>
        public InputList<Inputs.ServerPortGetArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.ServerPortGetArgs>());
            set => _ports = value;
        }

        /// <summary>
        /// restart policy of the server
        /// </summary>
        [Input("restartPolicy")]
        public Input<string>? RestartPolicy { get; set; }

        [Input("tags")]
        private InputList<Inputs.ServerTagGetArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.ServerTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ServerTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// [NAME-VERSION] the version of the hub server
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ServerState()
        {
        }
        public static new ServerState Empty => new ServerState();
    }
}
